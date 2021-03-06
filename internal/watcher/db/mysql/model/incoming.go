package model

import (
	"fmt"
	"math/big"
	"time"

	"github.com/void616/gm.mint.sender/internal/watcher/db/types"
	mint "github.com/void616/gm.mint"
	"github.com/void616/gm.mint/amount"
)

// Incoming model
type Incoming struct {
	ID            uint64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT:true;NOT NULL"`
	ServiceID     uint64 `gorm:"NOT NULL"`
	Service       Service
	To            []byte     `gorm:"SIZE:32;NOT NULL"`
	From          []byte     `gorm:"SIZE:32;NOT NULL"`
	Amount        string     `gorm:"NOT NULL" sql:"TYPE:decimal(30,18)"`
	Token         uint16     `gorm:"NOT NULL"`
	Digest        []byte     `gorm:"SIZE:32;NOT NULL"`
	Block         []byte     `gorm:"SIZE:32;NOT NULL"`
	Timestamp     time.Time  `gorm:"NOT NULL;DEFAULT:current_timestamp"`
	FirstNotifyAt *time.Time `gorm:""`
	NotifyAt      *time.Time `gorm:""`
	Notified      bool       `gorm:"NOT NULL"`
}

// MapFrom mapping
func (i *Incoming) MapFrom(t *types.Incoming) error {
	svc := Service{}
	if err := (&svc).MapFrom(&t.Service); err != nil {
		return err
	}
	i.ID = t.ID
	i.Service = svc
	i.To = t.To.Bytes()
	i.From = t.From.Bytes()
	i.Amount = t.Amount.String()
	i.Token = uint16(t.Token)
	i.Digest = t.Digest.Bytes()
	i.Block = t.Block.Bytes()
	i.Timestamp = t.Timestamp
	i.FirstNotifyAt = t.FirstNotifyAt
	i.NotifyAt = t.NotifyAt
	i.Notified = t.Notified
	return nil
}

// MapTo mapping
func (i *Incoming) MapTo() (*types.Incoming, error) {
	svc, err := (&i.Service).MapTo()
	if err != nil {
		return nil, err
	}
	to, err := mint.BytesToPublicKey(i.To)
	if err != nil {
		return nil, fmt.Errorf("invalid to")
	}
	from, err := mint.BytesToPublicKey(i.From)
	if err != nil {
		return nil, fmt.Errorf("invalid from")
	}
	amo, err := amount.FromString(i.Amount)
	if err != nil {
		return nil, fmt.Errorf("invalid amount")
	}
	digest, err := mint.BytesToDigest(i.Digest)
	if err != nil {
		return nil, fmt.Errorf("invalid digest")
	}
	block := new(big.Int).SetBytes(i.Block)

	return &types.Incoming{
		ID:            i.ID,
		Service:       *svc,
		To:            to,
		From:          from,
		Amount:        amo,
		Token:         mint.Token(i.Token),
		Digest:        digest,
		Block:         block,
		Timestamp:     i.Timestamp,
		FirstNotifyAt: i.FirstNotifyAt,
		NotifyAt:      i.NotifyAt,
		Notified:      i.Notified,
	}, nil
}
