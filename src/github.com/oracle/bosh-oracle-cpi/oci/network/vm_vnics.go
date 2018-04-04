package network

import (
	"fmt"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/oci"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"oracle/oci/core/client/compute"
	"oracle/oci/core/client/virtual_network"
	"oracle/oci/core/models"
)

const networkLogTag = "OCINetwork"

func FindVnicsAttachedToInstance(connector client.Connector, logger boshlog.Logger, instanceID string, compartmentId string) ([]*models.Vnic, error) {

	// Find all VnicAttachments associated with the given instance
	p := compute.NewListVnicAttachmentsParams()
	p.WithInstanceID(&instanceID).WithCompartmentID(compartmentId)
	r, err := connector.CoreSevice().Compute.ListVnicAttachments(p)
	if err != nil {
		return nil, fmt.Errorf("Error finding VnicAttachments for instance %s, %v",
			instanceID, oci.CoreModelErrorMsg(err))
	}

	vnics := []*models.Vnic{}
	attachmentHandler := func(attachmentID string, vnicID string) {
		req := virtual_network.NewGetVnicParams().WithVnicID(vnicID)
		res, err := connector.CoreSevice().VirtualNetwork.GetVnic(req)
		if err != nil {
			logger.Info(networkLogTag, " Encountered error %v while finding vnic %s for attachment %s. Skipping", err, attachmentID, vnicID)
		} else {
			vnics = append(vnics, res.Payload)
		}
	}

	for _, attachment := range r.Payload {
		waiter := NewVnicAttachmentWaiter(connector, logger, attachmentHandler)
		if err := waiter.WaitFor(*attachment.ID); err != nil {
			logger.Info(networkLogTag, " Encountered error %v while waiting for vnic attachment %s to complete. Skipping", err, *attachment.ID)
		}
	}
	return vnics, nil
}
