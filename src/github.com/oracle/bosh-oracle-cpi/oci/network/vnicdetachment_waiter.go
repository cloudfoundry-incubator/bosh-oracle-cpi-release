package network

import (
	"errors"
	"fmt"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshretry "github.com/cloudfoundry/bosh-utils/retrystrategy"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"oracle/oci/core/client/compute"
	"time"
)

type DetachedStateHandler func(attachmentID string, vnicID string)

type VnicDetachmentWaiter interface {
	WaitFor(attachmentID string) error
}

func NewVnicDetachmentWaiter(c client.Connector, l boshlog.Logger, h DetachedStateHandler) VnicDetachmentWaiter {
	return &vnicDetachmentWaiter{connector: c, logger: l, detachedHandler: h}
}

type vnicDetachmentWaiter struct {
	connector client.Connector
	logger    boshlog.Logger

	detachedHandler DetachedStateHandler
}

func (w *vnicDetachmentWaiter) WaitFor(attachmentID string) (err error) {

	getAttachmentState := func() (bool, error) {

		p := compute.NewGetVnicAttachmentParams().WithVnicAttachmentID(attachmentID)
		r, err := w.connector.CoreSevice().Compute.GetVnicAttachment(p)
		if err != nil {
			return false, err
		}
		switch *r.Payload.LifecycleState {
		case "DETACHING", "ATTACHED":
			return true, errors.New("Waiting")
		case "ATTACHING":
			return false, fmt.Errorf("Vnic Attachment %s unexpectedly attaching", attachmentID)
		case "DETACHED":
			if w.detachedHandler != nil {
				w.detachedHandler(attachmentID, r.Payload.VnicID)
			}
			return true, nil
		}
		return false, errors.New("Unknown state")
	}

	retryable := boshretry.NewRetryable(getAttachmentState)
	retryStrategy := boshretry.NewAttemptRetryStrategy(100, 1*time.Second, retryable, w.logger)

	w.logger.Debug(networkLogTag, "Waiting for VNIC attachment %s to be detached...", attachmentID)
	if err := retryStrategy.Try(); err != nil {
		w.logger.Debug(networkLogTag, "Error waiting to reach desired state %v. Giving up.", err)
		return err
	}
	w.logger.Debug(networkLogTag, "Detached VNIC attachment %s", attachmentID)
	return nil
}
