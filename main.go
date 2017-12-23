package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocms-io/gocms/context"
	"github.com/gocms-io/gocms/utility/log"
	"github.com/gocms-io/plugin-event-logger/init/database"
	"github.com/gocms-io/plugin-event-logger/init/repository"
	"github.com/gocms-io/plugin-event-logger/init/service"
	"github.com/gocms-io/plugin-event-logger/init/controller"
)

//go:generate apidoc -c ./ -i ./domain -o ./content/docs/ -f ".*\\.go$" -f ".*\\.js$"
func main() {

	port := flag.Int("port", 30001, "port to run on.")
	flag.Parse()

	// init database
	database := database.DefaultSQL()

	// migrate cms db
	database.SQL.MigrateSql()

	// start gin with defaults

	switch context.Config.EnvVars.LogLevel {
	case log.LOG_LEVEL_CRITICAL:
		fallthrough
	case log.LOG_LEVEL_ERROR:
		gin.SetMode(gin.ReleaseMode)
	case log.LOG_LEVEL_WARNING:
		gin.SetMode(gin.TestMode)
	case log.LOG_LEVEL_DEBUG:
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	r.Delims("{[{", "}]}")

	// setup repositories
	repositoriesGroup := repository.DefaultRepositoriesGroup(database.SQL.Dbx)

	// setup services
	servicesGroup := service.DefaultServicesGroup(repositoriesGroup, database)

	// setup controllers
	_ = controller.DefaultControllerGroup(r, servicesGroup)

	log.Infof("Listening on port: %v\n", *port)

	r.Run(fmt.Sprintf("localhost:%v", *port))

}
