package vm

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	"fmt"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"github.com/oracle/bosh-oracle-cpi/oci/network"
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
	"oracle/oci/core/client/virtual_network"
	"oracle/oci/core/models"
)

type vipNetworkConfigurator struct {
	connector client.Connector
	logger    boshlog.Logger

	configuration NetworkConfiguration
	vcnId         string
	subnetId      string
}

func NewVipNetworkConfigurator(c client.Connector, l boshlog.Logger, n NetworkConfiguration, vcnId string, subnetId string) VnicConfigurator {
	return &vipNetworkConfigurator{connector: c, logger: l, configuration: n, vcnId: vcnId, subnetId: subnetId}
}

func (vc vipNetworkConfigurator) CreatePrimaryVnicDetail(vnicName string) (models.CreateVnicDetails, error) {

	FALSE := false
	return models.CreateVnicDetails{
		AssignPublicIP: &FALSE,
		SubnetID:       &vc.subnetId,
		DisplayName:    vnicName}, nil
}

func (vc vipNetworkConfigurator) CreateSecondaryVnicDetail(vnicName string) (models.CreateVnicDetails, error) {

	FALSE := false
	return models.CreateVnicDetails{
		AssignPublicIP: &FALSE,
		SubnetID:       &vc.subnetId,
		DisplayName:    vnicName}, nil
}

func (vc vipNetworkConfigurator) ConfigurePrimaryVnic(in *resource.Instance) error {

	vc.logger.Debug(logTag, "Configuring primary vnic for instance %v", in.ID())

	in.WaitUntilStarted(vc.connector, vc.logger)

	privateIP, err := vc.primaryPrivateIpForInstance(in.ID())
	if err != nil {
		return err
	}

	publicIP, err := vc.configuredPublicIP()
	if err != nil {
		return err
	}

	// Associate public IP -> private IP
	d := models.UpdatePublicIPDetails{DisplayName: "vnc-primary", PrivateIPID: &privateIP.ID}
	p := virtual_network.NewUpdatePublicIPParams().WithPublicIPID(publicIP.ID).WithUpdatePublicIPDetails(&d)
	res, err := vc.connector.CoreSevice().VirtualNetwork.UpdatePublicIP(p)

	if err == nil {
		vc.logger.Debug(logTag, "Assigned public IP %+v", *res.Payload)
	}
	return err
}

func (vc vipNetworkConfigurator) configuredPublicIP() (*models.PublicIP, error) {

	details := models.GetPublicIPByIPAddressDetails{IPAddress: &vc.configuration.IP}
	p := virtual_network.NewGetPublicIPByIPAddressParams().WithGetPublicIPByIPAddressDetails(&details)
	res, err := vc.connector.CoreSevice().VirtualNetwork.GetPublicIPByIPAddress(p)

	if err != nil {
		return nil, err
	}
	publicIP := res.Payload
	if publicIP.LifecycleState != models.PublicIPLifecycleStateAVAILABLE {
		return nil, fmt.Errorf("Unexpected lifecycle state %s for public IP %s ", publicIP.LifecycleState, vc.configuration.IP)
	}
	return publicIP, nil
}

func (vc vipNetworkConfigurator) primaryPrivateIpForInstance(instanceID string) (*models.PrivateIP, error) {
	vnic, err := vc.primaryVnic(instanceID)
	if err != nil {
		return nil, err
	}
	return vc.primaryPrivateIP(*vnic.ID)
}

func (vc vipNetworkConfigurator) primaryVnic(instanceID string) (*models.Vnic, error) {

	vnics, err := network.FindVnicsAttachedToInstance(vc.connector, instanceID, vc.connector.CompartmentId())
	if err != nil {
		return nil, err
	}

	for _, v := range vnics {
		if v.IsPrimary {
			return v, nil
		}
	}
	return nil, fmt.Errorf("Unable to find primary vnic for instance %v", instanceID)
}

func (vc vipNetworkConfigurator) primaryPrivateIP(vnicID string) (*models.PrivateIP, error) {

	p := virtual_network.NewListPrivateIpsParams().WithVnicID(&vnicID)
	res, err := vc.connector.CoreSevice().VirtualNetwork.ListPrivateIps(p)

	if err != nil {
		return nil, err
	}
	for _, ip := range res.Payload {
		if ip.IsPrimary {
			return ip, nil
		}
	}
	return nil, fmt.Errorf("Unable to find primary private IP for vnic %v", vnicID)
}
