package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/gocms-io/plugin-event-logger/domain/setting/setting_repository"
	"github.com/gocms-io/plugin-event-logger/domain/event_logger/event_logger_repository"
)

type RepositoriesGroup struct {
	SettingsRepository    setting_repository.ISettingsRepository
	EventLoggerRepository event_logger_repository.IEventLoggerRepository
	dbx                   *sqlx.DB
}

func DefaultRepositoriesGroup(dbx *sqlx.DB) *RepositoriesGroup {

	// setup repositories
	rg := RepositoriesGroup{
		dbx:                   dbx,
		SettingsRepository:    setting_repository.DefaultSettingsRepository(dbx),
		EventLoggerRepository: event_logger_repository.DefaultEventLoggerRepository(dbx),
	}
	return &rg
}