package network

import (
	"fmt"
	"github.com/oracle/bosh-oracle-cpi/oci"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"oracle/oci/core/client/compute"
	"oracle/oci/core/client/virtual_network"
	"oracle/oci/core/models"
)

const networkLogTag = "OCINetwork"

func FindVnicsAttachedToInstance(connector client.Connector, instanceID string, compartmentId string) ([]*models.Vnic, error) {

	// Find all VnicAttachments associated with the given instance
	p := compute.NewListVnicAttachmentsParams()
	p.WithInstanceID(&instanceID).WithCompartmentID(compartmentId)
	r, err := connector.CoreSevice().Compute.ListVnicAttachments(p)
	if err != nil {
		return nil, fmt.Errorf("Error finding VnicAttachments for instance %s, %v",
			instanceID, oci.CoreModelErrorMsg(err))
	}

	vnics := []*models.Vnic{}
	for _, attachment := range r.Payload {

		switch *attachment.LifecycleState {
		case models.VnicAttachmentLifecycleStateATTACHED:
			req := virtual_network.NewGetVnicParams().WithVnicID(attachment.VnicID)
			res, err := connector.CoreSevice().VirtualNetwork.GetVnic(req)
			if err != nil {
				return nil, fmt.Errorf("Error finding Vnic for attachment %s. Reason:%s",
					*attachment.ID, oci.CoreModelErrorMsg(err))
			}
			vnics = append(vnics, res.Payload)

		case models.VnicAttachmentLifecycleStateATTACHING:
		case models.VnicAttachmentLifecycleStateDETACHED, models.VnicAttachmentLifecycleStateDETACHING:
		}
	}
	return vnics, nil
}
