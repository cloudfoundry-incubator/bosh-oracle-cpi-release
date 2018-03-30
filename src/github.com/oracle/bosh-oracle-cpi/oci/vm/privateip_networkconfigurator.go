package vm

import (
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
	"oracle/oci/core/models"
)

type privateIPConfigurator struct {
	configuration NetworkConfiguration
	vcnId         string
	subnetId      string
}

func NewPrivateIPConfigurator(n NetworkConfiguration, vcnId string, subnetId string) VnicConfigurator {
	return &privateIPConfigurator{configuration: n, vcnId: vcnId, subnetId: subnetId}
}

func (c privateIPConfigurator) CreatePrimaryVnicDetail(vnicName string) (models.CreateVnicDetails, error) {

	return models.CreateVnicDetails{
		PrivateIP:   c.configuration.IP,
		SubnetID:    &c.subnetId,
		DisplayName: vnicName}, nil
}

func (c privateIPConfigurator) CreateSecondaryVnicDetail(vnicName string) (models.CreateVnicDetails, error) {

	return models.CreateVnicDetails{
		PrivateIP:   c.configuration.IP,
		SubnetID:    &c.subnetId,
		DisplayName: vnicName}, nil
}

func (c privateIPConfigurator) ConfigurePrimaryVnic(_ *resource.Instance) error {
	return nil
}
