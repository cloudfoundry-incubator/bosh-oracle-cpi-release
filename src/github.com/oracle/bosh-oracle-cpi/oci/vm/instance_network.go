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
	VcnID      string
	SubnetName string
	SubnetID   string
	IP         string
	Type       string
}

func (n NetworkConfiguration) newVnicConfigurator(connector client.Connector, logger boshlog.Logger) (VnicConfigurator, error) {

	vcnID, subnetID, err := (&n).networkIDs(connector)
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

func (n *NetworkConfiguration) subnetID(connector client.Connector, vcnID string) (string, error) {

	if n.SubnetID != "" {
		return n.SubnetID, nil
	}
	_, err := n.vcnID(connector)
	if err != nil {
		return "", err
	}

	p := virtual_network.NewListSubnetsParams()
	p.WithCompartmentID(connector.CompartmentId()).WithVcnID(vcnID)
	response, err := connector.CoreSevice().VirtualNetwork.ListSubnets(p)
	if err != nil {
		return "", err
	}
	matches := 0
	var id string
	for _, s := range response.Payload {
		if s.DisplayName == n.SubnetName {
			matches += 1
			if matches > 1 {
				return "", fmt.Errorf("More than 1 subnet named '%s' found in vcn '%s'", n.VcnName, vcnID)
			}
			id = *s.ID
		}
	}
	if matches == 1 {
		n.SubnetID = id
		return id, nil
	}
	return "", fmt.Errorf("Unable to find OCID of subnet '%s' in vcn '%s'", n.SubnetName, vcnID)
}

// VcnID queries the OCID of a vcn from the compute service
func (n *NetworkConfiguration) vcnID(connector client.Connector) (string, error) {

	if n.VcnID != "" {
		return n.VcnID, nil
	}
	req := virtual_network.NewListVcnsParams()
	req.WithCompartmentID(connector.CompartmentId())
	res, err := connector.CoreSevice().VirtualNetwork.ListVcns(req)
	if err != nil {
		return "", err
	}

	matches := 0
	var id string
	for _, v := range res.Payload {
		if v.DisplayName == n.VcnName {
			matches += 1
			if matches > 1 {
				return "", fmt.Errorf("More than 1 vcn with named '%s' found in compartment '%s'", n.VcnName, connector.CompartmentId())
			}
			id = *v.ID
		}
	}
	if matches == 1 {
		n.VcnID = id
		return id, nil
	}
	return "", fmt.Errorf("Unable to finding OCID of VCN '%s' in compartment '%s'", n.VcnName, connector.CompartmentId())
}

func (n *NetworkConfiguration) networkIDs(connector client.Connector) (string, string, error) {

	vcnID, err := n.vcnID(connector)

	if err != nil {
		return "", "", err
	}

	subnetID, err := n.subnetID(connector, vcnID)

	return vcnID, subnetID, err
}

func (n NetworkConfiguration) validate() error {

	if n.VcnName == "" && n.VcnID == "" {
		return fmt.Errorf("Must specify either vcn name or vcn id")
	}
	if n.SubnetName == "" && n.SubnetID == "" {
		return fmt.Errorf("Must specify either subnet name or subnet id")
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
