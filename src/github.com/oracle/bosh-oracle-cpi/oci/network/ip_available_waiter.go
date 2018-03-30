package network

import (
	"errors"
	"fmt"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshretry "github.com/cloudfoundry/bosh-utils/retrystrategy"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"oracle/oci/core/client/virtual_network"
	"oracle/oci/core/models"
	"time"
)

// PublicIPWaiter allows callers to wait until a public IP reaches
// a desired state
type PublicIPWaiter interface {
	WaitUntilAvailable(ip *models.PublicIP) error
	WaitUntilAssigned(ip *models.PublicIP) error
}

type publicIPWaiter struct {
	connector client.Connector
	logger    boshlog.Logger
}

func NewPublicIPWaiter(c client.Connector, l boshlog.Logger) PublicIPWaiter {
	return &publicIPWaiter{connector: c, logger: l}
}

func (iw *publicIPWaiter) WaitUntilAvailable(ip *models.PublicIP) error {

	getIPState := func() (bool, error) {

		switch ip.LifecycleState {
		case models.PublicIPLifecycleStateASSIGNED,
			models.PublicIPLifecycleStateASSIGNING,
			models.PublicIPLifecycleStatePROVISIONING,
			models.PublicIPLifecycleStateUNASSIGNING,
			models.PublicIPLifecycleStateUNASSIGNED:
			var err error
			ip, err = queryIP(iw.connector, ip.ID)
			if err != nil {
				return false, err
			}
			return true, errors.New("Waiting")

		case models.PublicIPLifecycleStateAVAILABLE:
			return false, nil

		case models.PublicIPLifecycleStateTERMINATED:
			return false, fmt.Errorf("IP %s is terminated", ip.ID)
		default:
			return false, fmt.Errorf("Unknown ip lifecycle state %s", ip.LifecycleState)

		}
	}

	retryable := boshretry.NewRetryable(getIPState)
	delay := 1 * time.Second
	retryStrategy := boshretry.NewUnlimitedRetryStrategy(delay, retryable, iw.logger)

	iw.logger.Debug(networkLogTag, "Waiting for public IP to be available. Will check every %d secs", int(delay.Seconds()))
	if err := retryStrategy.Try(); err != nil {
		iw.logger.Debug(networkLogTag, "Error waiting %v", err)
		return err
	}
	iw.logger.Debug(networkLogTag, "Done")
	return nil

}

func queryIP(c client.Connector, id string) (*models.PublicIP, error) {

	p := virtual_network.NewGetPublicIPParams().WithPublicIPID(id)
	res, err := c.CoreSevice().VirtualNetwork.GetPublicIP(p)
	if err != nil {
		return nil, fmt.Errorf("Error finding public IP for %s", id)
	}
	return res.Payload, nil
}

func (iw *publicIPWaiter) WaitUntilAssigned(ip *models.PublicIP) error {

	getIPState := func() (bool, error) {

		switch ip.LifecycleState {
		case models.PublicIPLifecycleStateAVAILABLE,
			models.PublicIPLifecycleStateASSIGNING,
			models.PublicIPLifecycleStatePROVISIONING,
			models.PublicIPLifecycleStateUNASSIGNING,
			models.PublicIPLifecycleStateUNASSIGNED:
			var err error
			ip, err = queryIP(iw.connector, ip.ID)
			if err != nil {
				return false, err
			}
			return true, errors.New("Waiting")

		case models.PublicIPLifecycleStateASSIGNED:
			return false, nil

		case models.PublicIPLifecycleStateTERMINATED:
			return false, fmt.Errorf("IP %s is terminated", ip.ID)
		default:
			return false, fmt.Errorf("Unknown ip lifecycle state %s", ip.LifecycleState)

		}
	}

	retryable := boshretry.NewRetryable(getIPState)
	delay := 1 * time.Second
	retryStrategy := boshretry.NewUnlimitedRetryStrategy(delay, retryable, iw.logger)

	iw.logger.Debug(networkLogTag, "Waiting for public IP to be assigned. Will check every %d secs", int(delay.Seconds()))
	if err := retryStrategy.Try(); err != nil {
		iw.logger.Debug(networkLogTag, "Error waiting %v", err)
		return err
	}
	iw.logger.Debug(networkLogTag, "Done")
	return nil

}