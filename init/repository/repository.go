package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/gocms-io/plugin-event-logger/domain/setting/setting_repository"
)

type RepositoriesGroup struct {
	SettingsRepository    setting_repository.ISettingsRepository
	//EventLoggerRepository merchant_account_respository.IMerchantAccountRepository
	dbx                   *sqlx.DB
}

func DefaultRepositoriesGroup(dbx *sqlx.DB) *RepositoriesGroup {

	// setup repositories
	rg := RepositoriesGroup{
		dbx:                   dbx,
		SettingsRepository:    setting_repository.DefaultSettingsRepository(dbx),
		//MerchantAccountRepository: merchant_account_respository.DefaultMerchantAccountRepository(dbx),
	}
	return &rg
}