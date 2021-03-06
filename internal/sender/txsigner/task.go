package txsigner

import (
	"math/big"
	"time"

	"github.com/void616/gm.mint.rpc/request"
	"github.com/void616/gm.mint.sender/internal/sender/db/types"
	"github.com/void616/gotask"
)

// Task loop
func (s *Signer) Task(token *gotask.Token) {

	requests := make(chan interface{}, itemsPerShot*4)
	defer close(requests)

	currentBlock := new(big.Int)

	for !token.Stopped() {

		// get current network block
		{
			ctx, conn, cls, err := s.pool.Conn()
			if err != nil {
				s.logger.WithError(err).Error("Failed to get RPC connection")
				token.Sleep(time.Second * 30)
				continue
			}

			state, rerr, err := request.GetBlockchainState(ctx, conn)
			if err != nil || rerr != nil {
				cls()
				if rerr != nil {
					err = rerr.Err()
				}
				s.logger.WithError(err).Error("Failed to get current block ID")
				token.Sleep(time.Second * 30)
				continue
			}

			currentBlock.Sub(state.BlockCount.Int, big.NewInt(1))
			cls()
		}

		count := 0

		// get stale requests (approvements)
		{
			elderThan := new(big.Int).Sub(currentBlock, big.NewInt(staleAfterBlocks))

			list, err := s.dao.ListStaleApprovements(elderThan, itemsPerShot)
			if err != nil {
				s.logger.WithError(err).Error("Failed to get stale transactions")
				token.Sleep(time.Second * 30)
				continue
			}
			for _, v := range list {
				select {
				case requests <- v:
				default:
				}
			}
			count += len(list)
		}

		// get new requests (approvements)
		{
			list, err := s.dao.ListEnqueuedApprovements(itemsPerShot)
			if err != nil {
				s.logger.WithError(err).Error("Failed to get new transactions")
				token.Sleep(time.Second * 30)
				continue
			}
			for _, v := range list {
				select {
				case requests <- v:
				default:
				}
			}
			count += len(list)
		}

		// get stale requests (sendings)
		{
			elderThan := new(big.Int).Sub(currentBlock, big.NewInt(staleAfterBlocks))

			list, err := s.dao.ListStaleSendings(elderThan, itemsPerShot)
			if err != nil {
				s.logger.WithError(err).Error("Failed to get stale transactions")
				token.Sleep(time.Second * 30)
				continue
			}
			for _, v := range list {
				select {
				case requests <- v:
				default:
				}
			}
			count += len(list)
		}

		// get new requests (sendings)
		{
			list, err := s.dao.ListEnqueuedSendings(itemsPerShot)
			if err != nil {
				s.logger.WithError(err).Error("Failed to get new transactions")
				token.Sleep(time.Second * 30)
				continue
			}
			for _, v := range list {
				select {
				case requests <- v:
				default:
				}
			}
			count += len(list)
		}

		// empty queue
		if count == 0 {
			token.Sleep(time.Second * 10)
			continue
		}

		// metrics
		if s.metrics != nil {
			s.metrics.Queue.Set(float64(count))
		}

		// process queue
		processed := 0
		out := false
		for !out {
			select {
			default:
				out = true
			case item := <-requests:
				switch m := item.(type) {
				case *types.Approvement:
					if s.processApprovingRequest(m, currentBlock) {
						processed++
					}
				case *types.Sending:
					if s.processSendingRequest(m, currentBlock) {
						processed++
					}
				}
			}
		}

		if s.metrics != nil {
			s.metrics.Queue.Set(0)
		}

		if processed == 0 {
			token.Sleep(time.Second * 30)
		}
	}
}
