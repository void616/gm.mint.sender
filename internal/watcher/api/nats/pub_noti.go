package nats

import (
	"fmt"
	"time"

	proto "github.com/golang/protobuf/proto"
	walletsvc "github.com/void616/gm-mint-sender/pkg/watcher/nats/wallet"
	sumuslib "github.com/void616/gm-sumuslib"
	"github.com/void616/gm-sumuslib/amount"
)

// NotifyRefilling sends an event
func (n *Nats) NotifyRefilling(service string, to, from sumuslib.PublicKey, t sumuslib.Token, a *amount.Amount, tx sumuslib.Digest) error {

	// metrics
	if n.metrics != nil {
		defer func(t time.Time) {
			n.metrics.NotificationDuration.Observe(time.Since(t).Seconds())
		}(time.Now())
	}

	reqModel := &walletsvc.RefillEvent{
		Service:     service,
		PublicKey:   to.String(),
		From:        from.String(),
		Token:       t.String(),
		Amount:      a.String(),
		Transaction: tx.String(),
	}

	req, err := proto.Marshal(reqModel)
	if err != nil {
		return err
	}

	msg, err := n.natsConnection.Request(n.subjPrefix+walletsvc.SubjectRefill, req, time.Second*5)
	if err != nil {
		return err
	}

	repModel := walletsvc.RefillEventReply{}
	if err := proto.Unmarshal(msg.Data, &repModel); err != nil {
		return err
	}

	if !repModel.GetSuccess() {
		return fmt.Errorf("service rejection: %v", repModel.GetError())
	}

	return nil
}