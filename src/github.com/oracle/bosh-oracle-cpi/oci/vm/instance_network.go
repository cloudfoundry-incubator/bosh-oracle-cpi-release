package vm

import (
	"fmt"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"oracle/oci/core/client/virtual_network"
)

const (
	manualNetwork  = "manual"
	vipNetwork     = "vip"
	dynamicNetwork = "dynamic"
)

type NetworkConfiguration struct {
	VcnName    string
	SubnetName string
	IP         string
	Type       string
}

func (n NetworkConfiguration) newVnicConfigurator(connector client.Connector, logger boshlog.Logger) (VnicConfigurator, error) {

	vcnID, subnetID, err := n.networkIDs(connector)
	if err != nil {
		return nil, err
	}

	switch {
	case n.isManual(), n.isDynamic():
		return NewPrivateIPConfigurator(n, vcnID, subnetID), nil
	case n.isVip():
		return NewPublicIPConfigurator(connector, logger, n, vcnID, subnetID), nil
	}
	return nil, fmt.Errorf("Unsupported network type %s", n.Type)
}

func (n NetworkConfiguration) subnetID(connector client.Connector, vcnId string) (string, error) {

	_, err := n.vcnID(connector)
	if err != nil {
		return "", err
	}

	p := virtual_network.NewListSubnetsParams()
	p.WithCompartmentID(connector.CompartmentId()).WithVcnID(vcnId)
	response, err := connector.CoreSevice().VirtualNetwork.ListSubnets(p)
	if err != nil {
		return "", err
	}
	for _, s := range response.Payload {
		if s.DisplayName == n.SubnetName {
			return *s.ID, nil
		}
	}
	return "", fmt.Errorf("Unable to find ID of subnet %s", n.SubnetName)
}

// VcnID queries the OCID of a vcn from the compute service
func (n NetworkConfiguration) vcnID(connector client.Connector) (string, error) {

	req := virtual_network.NewListVcnsParams()
	req.WithCompartmentID(connector.CompartmentId())
	res, err := connector.CoreSevice().VirtualNetwork.ListVcns(req)
	if err != nil {
		return "", err
	}

	for _, v := range res.Payload {
		if v.DisplayName == n.VcnName {
			return *v.ID, nil
		}
	}
	return "", fmt.Errorf("Error finding VcnID of VCN %s", n.VcnName)
}

func (n NetworkConfiguration) networkIDs(connector client.Connector) (string, string, error) {

	vcnID, err := n.vcnID(connector)

	if err != nil {
		return "", "", err
	}

	subnetID, err := n.subnetID(connector, vcnID)

	return vcnID, subnetID, err
}

func (n NetworkConfiguration) validate() error {

	if n.VcnName == "" {
		return fmt.Errorf(" Missing VCN name")
	}
	if n.SubnetName == "" {
		return fmt.Errorf("Missing subnet name")
	}
	return nil
}

// isDynamic returns true if the network is configured
// as a "dynamic" network
func (n NetworkConfiguration) isDynamic() bool {
	return n.Type == dynamicNetwork
}

// isStatic returns true if the network is configured
// as a "manual" network
func (n NetworkConfiguration) isManual() bool {
	return n.Type == "" || n.Type == manualNetwork
}

// isVip returns true if the network is configured
// as a "vip" network
func (n NetworkConfiguration) isVip() bool {
	return n.Type == vipNetwork
}
