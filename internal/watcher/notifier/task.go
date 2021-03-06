package notifier

import (
	"time"

	"github.com/void616/gm.mint.sender/internal/watcher/db/types"
	"github.com/void616/gotask"
)

// Task loop
func (n *Notifier) Task(token *gotask.Token) {

	for !token.Stopped() {

		// get list
		list, err := n.dao.ListUnnotifiedIncomings(itemsPerShot)
		if err != nil {
			n.logger.WithError(err).Error("Failed to get unsent items")
			token.Sleep(time.Second * 30)
			continue
		}

		// nothing
		if len(list) == 0 {
			token.Sleep(time.Second * 30)
			continue
		}

		out := false
		for _, inc := range list {
			if out {
				break
			}

			// mark as notified
			{
				now := time.Now().UTC()
				if inc.FirstNotifyAt == nil {
					inc.FirstNotifyAt = &now
				}
				inc.NotifyAt = &now
				inc.Notified = true
			}
			if err := n.dao.UpdateIncoming(inc); err != nil {
				n.logger.
					WithError(err).
					WithField("wallet", inc.To.String()).
					WithField("tx", inc.Digest.String()).
					Error("Failed to update incoming")
				token.Sleep(time.Second * 30)
				out = true
				continue
			}

			// notify
			var notiErr error
			switch inc.Service.Transport {
			case types.ServiceNats:
				if n.natsTrans != nil {
					notiErr = n.natsTrans.NotifyRefilling(inc.Service.Name, inc.To, inc.From, inc.Token, inc.Amount, inc.Digest)
				} else {
					n.logger.Warn("Nats transport is disabled, skipping notification")
					continue
				}
			case types.ServiceHTTP:
				if n.httpTrans != nil {
					if inc.Service.CallbackURL != "" {
						notiErr = n.httpTrans.NotifyRefilling(inc.Service.CallbackURL, inc.Service.Name, inc.To, inc.From, inc.Token, inc.Amount, inc.Digest)
					}
				} else {
					n.logger.Warn("HTTP transport is disabled, skipping notification")
					continue
				}
			default:
				n.logger.Errorf("Transport %v is not implemented", inc.Service.Transport)
				continue
			}

			if notiErr != nil {
				n.logger.
					WithField("wallet", inc.To.String()).
					WithField("tx", inc.Digest.String()).
					WithError(notiErr).
					Error("Failed to notify")

				// notify next time
				when := time.Now().UTC()
				if inc.FirstNotifyAt != nil {
					mikes := time.Now().UTC().Sub(*inc.FirstNotifyAt).Minutes()
					switch {
					// for 5m: every 1m
					case mikes < 5:
						when = when.Add(time.Minute)
					// then for 30m: every 5m
					case mikes < 35:
						when = when.Add(time.Minute * 5)
					// then for 60m: every 10m
					case mikes < 95:
						when = when.Add(time.Minute * 10)
					// then every 120m
					default:
						when = when.Add(time.Minute * 120)
					}
				} else {
					when = when.Add(time.Hour * 24 * 365)
				}

				// mark as unnotified
				inc.NotifyAt = &when
				inc.Notified = false
				if err := n.dao.UpdateIncoming(inc); err != nil {
					n.logger.
						WithError(err).
						WithField("wallet", inc.To.String()).
						WithField("tx", inc.Digest.String()).
						Error("Failed to update incoming")
					token.Sleep(time.Second * 30)
					out = true
					continue
				}
			} else {
				n.logger.
					WithField("wallet", inc.To.String()).
					WithField("tx", inc.Digest.String()).
					Info("Notified")
			}
		}
	}
}
