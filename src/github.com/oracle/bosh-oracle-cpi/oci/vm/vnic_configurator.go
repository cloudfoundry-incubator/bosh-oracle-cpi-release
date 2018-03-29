package vm

import (
	"oracle/oci/core/models"
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
)

type VnicConfigurator interface {
	CreatePrimaryVnicDetail(vnicName string) (models.CreateVnicDetails, error)
	CreateSecondaryVnicDetail(vnicName string) (models.CreateVnicDetails, error)
	ConfigurePrimaryVnic(in *resource.Instance) error
}

type manualNetworkConfigurator struct {
	configuration NetworkConfiguration
	vcnId         string
	subnetId      string
}

func NewManualNetworkConfigurator(n NetworkConfiguration, vcnId string, subnetId string) VnicConfigurator {
	return &manualNetworkConfigurator{configuration: n, vcnId: vcnId, subnetId: subnetId}
}

func (c manualNetworkConfigurator) CreatePrimaryVnicDetail(vnicName string) (models.CreateVnicDetails, error) {

	return models.CreateVnicDetails{
		PrivateIP:   c.configuration.IP,
		SubnetID:    &c.subnetId,
		DisplayName: vnicName}, nil
}

func (c manualNetworkConfigurator) CreateSecondaryVnicDetail(vnicName string) (models.CreateVnicDetails, error) {

	return models.CreateVnicDetails{
		PrivateIP:   c.configuration.IP,
		SubnetID:    &c.subnetId,
		DisplayName: vnicName}, nil
}

func (c manualNetworkConfigurator) ConfigurePrimaryVnic(_ *resource.Instance) error {
	return nil
}
