package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/void616/gm.mint.sender/internal/sender/db/mysql/model"
	gormigrate "gopkg.in/gormigrate.v1"
)

var migrations = []*gormigrate.Migration{

	// initial
	{
		ID: "2019-09-26T13:20:00.350Z",
		Migrate: func(tx *gorm.DB) error {
			return tx.
				CreateTable(&model.Wallet{}).
				CreateTable(&model.Sending{}).
				AddUniqueIndex("ux_sender_sendings_service_requestid", "service", "request_id").
				AddIndex("ix_sender_sendings_status", "status").
				AddIndex("ix_sender_sendings_sentatblock", "sent_at_block").
				CreateTable(&model.Approvement{}).
				AddUniqueIndex("ux_sender_approvs_service_requestid", "service", "request_id").
				AddIndex("ix_sender_approvs_status", "status").
				AddIndex("ix_sender_approvs_sentatblock", "sent_at_block").
				Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.
				DropTable(&model.Approvement{}).
				DropTable(&model.Sending{}).
				DropTable(&model.Wallet{}).
				Error
		},
	},

	// sendings: optional wallet approvement ignoring
	{
		ID: "2020-02-21T20:12:48.866Z",
		Migrate: func(tx *gorm.DB) error {
			return tx.
				AutoMigrate(&model.Sending{}).
				Error
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	},
}
