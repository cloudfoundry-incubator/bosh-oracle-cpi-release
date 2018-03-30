package resource

import (
	"github.com/oracle/bosh-oracle-cpi/oci/client"

	"fmt"
	"github.com/oracle/bosh-oracle-cpi/oci/network"
	"oracle/oci/core/models"
)

type Location struct {
	availabilityDomain string
	compartmentId      string
}

func NewLocation(ad string, compartmentId string) Location {
	return Location{availabilityDomain: ad, compartmentId: compartmentId}
}

func (loc Location) instanceIPs(connector client.Connector, instanceID string) (
	publicip []string, privateip []string, err error) {

	vnics, err := loc.vnics(connector, instanceID)
	if err != nil {
		return nil, nil, err
	}
	public := make([]string, len(vnics))
	private := make([]string, len(vnics))
	for i, v := range vnics {
		public[i] = v.PublicIP
		private[i] = *v.PrivateIP
	}
	return public, private, nil
}

func (loc Location) vnics(connector client.Connector, instanceID string) ([]*models.Vnic, error) {

	vnics, err := network.FindVnicsAttachedToInstance(connector, instanceID, loc.compartmentId)
	if err != nil {
		return nil, err
	}
	if len(vnics) == 0 {
		err = fmt.Errorf("No Vnic Attachments found for VM %s", instanceID)
	}
	return vnics, err
}

func (loc Location) CompartmentID() string {
	return loc.compartmentId
}

func (loc Location) AvailabilityDomain() string {
	return loc.availabilityDomain
}
