package test

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"github.com/oracle/bosh-oracle-cpi/oci/network"
	"oracle/oci/core/client/virtual_network"
	"oracle/oci/core/models"
	"testing"
)

type PublicIPFixture struct {
	connectionFixture *ConnectionFixture
	publicIP          *models.PublicIP
}

func NewPublicIPFixture() *PublicIPFixture {
	return &PublicIPFixture{connectionFixture: NewConnectionFixture()}
}

func (p *PublicIPFixture) Setup(t *testing.T) error {
	p.connectionFixture.Setup(t)
	vn := p.connectionFixture.connector.CoreSevice().VirtualNetwork

	cid := p.connectionFixture.Connector().CompartmentId()
	kind := models.CreatePublicIPDetailsLifetimeRESERVED
	params := virtual_network.NewCreatePublicIPParams().WithCreatePublicIPDetails(&models.CreatePublicIPDetails{
		CompartmentID: &cid,
		DisplayName:   "public-ip-test-fixture",
		Lifetime:      &kind,
	})

	res, err := vn.CreatePublicIP(params)
	if err != nil {
		t.Fatal(err)
	}
	p.publicIP = res.Payload

	waiter := network.NewPublicIPWaiter(p.Connector(), p.Logger())
	if err := waiter.WaitUntilAvailable(p.publicIP); err != nil {
		t.Fatal(err)
	}
	return nil
}

func (p *PublicIPFixture) TearDown(t *testing.T) error {

	vn := p.connectionFixture.connector.CoreSevice().VirtualNetwork
	params := virtual_network.NewDeletePublicIPParams().WithPublicIPID(p.publicIP.ID)
	_, err := vn.DeletePublicIP(params)

	if err != nil {
		t.Logf(" Ignoring error while deleting IP %s", p.publicIP.ID, err)
	}
	return p.connectionFixture.TearDown(t)
}

func (p PublicIPFixture) Connector() client.Connector {
	return p.connectionFixture.Connector()
}
func (p PublicIPFixture) Logger() boshlog.Logger {
	return p.connectionFixture.Logger()
}

func (p PublicIPFixture) AD() string {
	return p.connectionFixture.AD()
}

func (p PublicIPFixture) Address() string {
	return p.publicIP.IPAddress
}

func (p PublicIPFixture) ConnectionFixture() *ConnectionFixture {
	return p.connectionFixture
}