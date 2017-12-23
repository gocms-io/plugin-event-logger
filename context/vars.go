package context

import (
	"log"
	"github.com/gocms-io/plugin-event-logger/domain/setting/setting_model"
	"strings"
)

type envVars struct {
	// DB (GET FROM ENV)
	DbName     string
	DbUser     string
	DbPassword string
	DbServer   string
}

type dbVars struct {
	// App config
	SettingsRefreshRate int64
	IgnoreHeaders []string
}

func (dbVars *dbVars) LoadDbVars(settings map[string]setting_model.Setting) {

	log.Printf("Refresh Event Logger Services Settings\n")

	// App config
	dbVars.SettingsRefreshRate = GetIntOrFail("SETTINGS_REFRESH_RATE", settings)

	// Headers To Ignore
	ignoreHeadersStr := GetStringOrFail("IGNORE_HEADERS", settings)
	dbVars.IgnoreHeaders = strings.Split(ignoreHeadersStr, ",")
}
