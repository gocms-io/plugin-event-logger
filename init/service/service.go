package service

import (
	"time"
	"github.com/gocms-io/plugin-event-logger/domain/setting/setting_service"
	"github.com/gocms-io/plugin-event-logger/domain/health/health_service"
	"github.com/gocms-io/plugin-event-logger/init/repository"
	"github.com/gocms-io/plugin-event-logger/init/database"
	"github.com/gocms-io/plugin-event-logger/context"
)

type ServicesGroup struct {
	SettingsService        setting_service.ISettingsService
	HealthService          health_service.IHealthService
	//MerchantAccountService merchant_account_service.IMerchantAccountService
	repositoriesGroup      *repository.RepositoriesGroup
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
	//merchantAccountService := merchant_account_service.DefaultMerchantAccountService(repositoriesGroup)

	sg := &ServicesGroup{
		SettingsService:        settingsService,
		HealthService:          healthService,
		//MerchantAccountService: merchantAccountService,
		repositoriesGroup:      repositoriesGroup,
	}
	return sg
}
