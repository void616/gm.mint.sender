package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	mint "github.com/void616/gm.mint"
	pkg "github.com/void616/gm.mint.sender/pkg/watcher/http"
	"github.com/void616/gm.mint/amount"
)

// NotifyRefilling sends a notification
func (h *HTTP) NotifyRefilling(url, service string, to, from mint.PublicKey, t mint.Token, a *amount.Amount, d mint.Digest) error {
	// metrics
	if h.metrics != nil {
		defer func(t time.Time) {
			h.metrics.NotificationDuration.Observe(time.Since(t).Seconds())
		}(time.Now())
	}

	event := pkg.RefillEvent{
		Service:     service,
		PublicKey:   to.String(),
		From:        from.String(),
		Token:       t.String(),
		Amount:      a.String(),
		Transaction: d.String(),
	}

	b, err := json.Marshal(&event)
	if err != nil {
		return err
	}

	// http request
	{
		timeoutSec := 10
		transport := &http.Transport{
			IdleConnTimeout: time.Second * time.Duration(timeoutSec),
		}
		client := &http.Client{
			Timeout:   time.Second * time.Duration(timeoutSec),
			Transport: transport,
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return fmt.Errorf("callback status code is %v", resp.StatusCode)
		}
	}

	return nil
}
