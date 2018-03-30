package vm

import (
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
	"oracle/oci/core/models"
)

type VnicConfigurator interface {
	CreatePrimaryVnicDetail(vnicName string) (models.CreateVnicDetails, error)
	CreateSecondaryVnicDetail(vnicName string) (models.CreateVnicDetails, error)
	ConfigurePrimaryVnic(in *resource.Instance) error
}
