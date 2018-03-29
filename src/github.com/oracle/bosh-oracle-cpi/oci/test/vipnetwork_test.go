package test

import (
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
	"github.com/oracle/bosh-oracle-cpi/oci/vm"
	"testing"
)

func Test_VMAssignFloatingIP(t *testing.T) {

	state := NewConnectionFixture()
	state.Setup(t)
	defer state.TearDown(t)

	// Creator and Terminator
	var in *resource.Instance
	var err error

	creator := vm.NewCreator(state.Connector(),
		state.Logger(), state.AD())
	terminator := vm.NewTerminator(state.Connector(), state.Logger())

	deleteInstance := func() {
		if err == nil && in != nil {
			terminator.TerminateInstance(in.ID())
			in = nil
		}
	}
	defer deleteInstance()

	// Create a VM
	icfg := state.DefaultInstanceConfiguration()
	icfg.Network = []vm.NetworkConfiguration{
		{VcnName: state.VCN(), SubnetName: state.Subnet(), IP: "129.146.151.194", Type: "vip"},
	}

	icfg.Name = "test-instance-with-floating-ip"
	in, err = creator.CreateInstance(icfg, vm.InstanceMetadata{})
	if err != nil {
		t.Fatalf("Error creating initial instance. Err: %v", err)
	}

	assignedPublicIP, err := in.PublicIP(state.Connector(), state.Logger())
	if err != nil {
		t.Fatalf("Error getting assigned public IP. Err: %v", err)
	}

	assertEqual(t, assignedPublicIP, "129.146.151.194", "")

}
