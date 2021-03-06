package vm

import (
	"fmt"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/oci"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"github.com/oracle/bosh-oracle-cpi/oci/network"
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
	"oracle/oci/core/client/compute"
	"oracle/oci/core/models"
)

const logTag = "OCIVMOperations"

type InstanceConfiguration struct {
	ImageId string
	Shape   string
	Name    string
	Network Networks
}

type Creator interface {
	CreateInstance(icfg InstanceConfiguration, md InstanceMetadata) (*resource.Instance, error)
}

type CreatorFactory func(client.Connector, boshlog.Logger, string) Creator

type creator struct {
	connector client.Connector
	logger    boshlog.Logger
	location  resource.Location
}

func NewCreator(c client.Connector, l boshlog.Logger, availabilityDomain string) Creator {
	return &creator{connector: c, logger: l,
		location:              resource.NewLocation(availabilityDomain, c.CompartmentId()),
	}
}

func (cv *creator) CreateInstance(icfg InstanceConfiguration,
	md InstanceMetadata) (*resource.Instance, error) {

	if err := icfg.Network.validate(); err != nil {
		return nil, err
	}
	return cv.launchInstance(icfg, md)
}
func (cv *creator) launchInstance(icfg InstanceConfiguration, md InstanceMetadata) (*resource.Instance, error) {

	// Arrange for primary VNIC
	configurator, err := icfg.Network.primary().newVnicConfigurator(cv.connector, cv.logger)
	if err != nil {
		return nil, newLaunchInstanceError(err)
	}
	primaryVnic, err := configurator.CreatePrimaryVnicDetail("primary")
	if err != nil {
		return nil, newLaunchInstanceError(err)
	}

	// Create Instance
	req := cv.buildLaunchInstanceParams(icfg, md, &primaryVnic)
	cv.logLaunchingInstanceDebugMsg(req)
	res, err := cv.connector.CoreSevice().Compute.LaunchInstance(req)
	if err != nil {
		return nil, newLaunchInstanceError(err)
	}
	instance := resource.NewInstance(*res.Payload.ID, cv.location)

	// Complete primary VNIC configuration
	if err := configurator.ConfigurePrimaryVnic(instance); err != nil {
		return nil, newLaunchInstanceError(err)
	}

	// Additional VNICs
	if icfg.Network.hasSecondaries() {
		err := cv.attachSecondaryVnics(instance, icfg.Network.secondaries())
		if err != nil {
			return nil, newLaunchInstanceError(err)
		}
	}
	return instance, nil
}

func (cv *creator) buildLaunchInstanceParams(icfg InstanceConfiguration, md InstanceMetadata,
	primaryVnic *models.CreateVnicDetails) *compute.LaunchInstanceParams {

	req := compute.NewLaunchInstanceParams()
	ad := cv.location.AvailabilityDomain()
	cid := cv.location.CompartmentID()

	details := &models.LaunchInstanceDetails{
		AvailabilityDomain: &ad,
		DisplayName:        icfg.Name,
		CompartmentID:      &cid,
		Shape:              &icfg.Shape,
		CreateVnicDetails:  primaryVnic,
		ImageID:            icfg.ImageId,
	}
	details.Metadata = md.AsMap()

	return req.WithLaunchInstanceDetails(details)
}

func extractMsgFromError(err error) string {
	return oci.CoreModelErrorMsg(err)
}

func newLaunchInstanceError(err error) error {
	return fmt.Errorf("Error launching instance. Reason: %s", extractMsgFromError(err))
}

func (cv *creator) logLaunchingInstanceDebugMsg(p *compute.LaunchInstanceParams) {

	fmtStr := "LaunchInstance: AD:%s, Name:%s, Shape:%s\nCompartmentId:%s\nImageId:%s\n"
	args := []interface{}{*p.LaunchInstanceDetails.AvailabilityDomain,
						  p.LaunchInstanceDetails.DisplayName,
						  *p.LaunchInstanceDetails.Shape,
						  *p.LaunchInstanceDetails.CompartmentID,
						  p.LaunchInstanceDetails.ImageID,
	}
	if p.LaunchInstanceDetails.CreateVnicDetails != nil {
		fmtStr += "Subnet:%s, PrivateIP:%s\n"
		args = append(args, *p.LaunchInstanceDetails.CreateVnicDetails.SubnetID,
			p.LaunchInstanceDetails.CreateVnicDetails.PrivateIP)

		if p.LaunchInstanceDetails.CreateVnicDetails.AssignPublicIP != nil {
			fmtStr += "AssignPublicIP:%v\n"
			args = append(args, *p.LaunchInstanceDetails.CreateVnicDetails.AssignPublicIP)
		}
	}
	fmtStr += "Metadata:\n\t%v\n"
	args = append(args, p.LaunchInstanceDetails.Metadata)

	cv.logger.Debug(logTag, fmtStr, args...)
}

func (cv *creator) attachSecondaryVnics(in *resource.Instance, secondaries []NetworkConfiguration) error {

	var attachmentError error
	deleteInstance := func() {
		if attachmentError != nil {
			NewTerminator(cv.connector, cv.logger).TerminateInstance(in.ID())
		}
	}
	defer deleteInstance()

	for i, secondary := range secondaries {
		c, err := secondary.newVnicConfigurator(cv.connector, cv.logger)
		if err != nil {
			return err
		}
		vnicDetail, _ := c.CreateSecondaryVnicDetail(fmt.Sprintf("secondary-%d", i))
		attachmentError = cv.attachVnic(in, vnicDetail)
		if attachmentError != nil {
			return attachmentError
		}
	}
	return nil
}

func (cv *creator) attachVnic(in *resource.Instance, details models.CreateVnicDetails) error {

	in.WaitUntilStarted(cv.connector, cv.logger)

	instanceID := in.ID()

	attachmentDetails := models.AttachVnicDetails{
		InstanceID:        &instanceID,
		CreateVnicDetails: &details,
	}
	p := compute.NewAttachVnicParams().WithAttachVnicDetails(&attachmentDetails)
	res, err := cv.connector.CoreSevice().Compute.AttachVnic(p)
	if err != nil {
		return err
	}

	waiter := network.NewVnicAttachmentWaiter(cv.connector, cv.logger,
		func(attachmentID string, vnicID string) {
			cv.logger.Debug(logTag, "Attached Vnic to Instance %s. VnicID=%s", in.ID(), vnicID)
		})
	return waiter.WaitFor(*res.Payload.ID)
}
