package health_service

import (
	"github.com/gocms-io/gocms/utility/log"
	"time"
	"github.com/gocms-io/plugin-event-logger/domain/health/health_model"
	"github.com/gocms-io/plugin-event-logger/context"
	"github.com/gocms-io/plugin-event-logger/init/database"
)

type IHealthService interface {
	GetHealthStatus() (ok bool, context []string)
}

type HealthService struct {
	db     *database.Database
	health *health_model.HealthMonitor
}

func DefaultHealthService(database *database.Database) *HealthService {

	healthService := &HealthService{
		db: database,
		// default good until bad
		health: &health_model.HealthMonitor{
			Database: true,
		},
	}

	// add health checks
	context.Schedule.AddTicker(15*time.Second, healthService.checkDatabaseHealth)

	return healthService

}

func (healthService *HealthService) GetHealthStatus() (ok bool, context []string) {
	// set ok until something is wrong
	ok = true

	// if database is not healthy
	if !healthService.health.Database {
		ok = false
		context = append(context, "Database connection lost")
	}

	return ok, context

}

func (healthService *HealthService) checkDatabaseHealth() {
	go func() {
		// check for database connectivity
		err := healthService.db.SQL.Dbx.Ping()
		if err != nil { // no connectivity
			log.Errorf("[Health Service] - Database connection lost: %v\n", err.Error())
			healthService.health.Database = false
		} else { // good connectivity
			healthService.health.Database = true
		}
	}()
}
