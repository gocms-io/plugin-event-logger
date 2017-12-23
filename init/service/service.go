package service

import (
	"time"
	"github.com/gocms-io/plugin-event-logger/domain/setting/setting_service"
	"github.com/gocms-io/plugin-event-logger/domain/health/health_service"
	"github.com/gocms-io/plugin-event-logger/init/repository"
	"github.com/gocms-io/plugin-event-logger/init/database"
	"github.com/gocms-io/plugin-event-logger/context"
	"github.com/gocms-io/plugin-event-logger/domain/event_logger/event_logger_service"
)

type ServicesGroup struct {
	SettingsService    setting_service.ISettingsService
	HealthService      health_service.IHealthService
	EventLoggerService event_logger_service.IEventLoggerService
	repositoriesGroup  *repository.RepositoriesGroup
}

func DefaultServicesGroup(repositoriesGroup *repository.RepositoriesGroup, db *database.Database) *ServicesGroup {

	// setup settings
	settingsService := setting_service.DefaultSettingsService(repositoriesGroup)
	settingsService.RegisterRefreshCallback(context.Config.DbVars.LoadDbVars)

	// refresh settings every x minutes
	refreshSettings := time.Duration(context.Config.DbVars.SettingsRefreshRate) * time.Minute
	context.Schedule.AddTicker(refreshSettings, func() {
		settingsService.RefreshSettingsCache()
	})

	// setup health service
	healthService := health_service.DefaultHealthService(db)
	eventLoggerService := event_logger_service.DefaultEventLoggerService(repositoriesGroup)

	sg := &ServicesGroup{
		SettingsService:    settingsService,
		HealthService:      healthService,
		EventLoggerService: eventLoggerService,
		repositoriesGroup:  repositoriesGroup,
	}
	return sg
}
