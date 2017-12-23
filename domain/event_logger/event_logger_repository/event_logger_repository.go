package event_logger_repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/gocms-io/plugin-event-logger/domain/event_logger/event_logger_model"
	"github.com/gocms-io/gocms/utility/log"
	"fmt"
)

type IEventLoggerRepository interface {
	Add(event *event_logger_model.Event) error
}

type EventLoggerRepository struct {
	database *sqlx.DB
}

func DefaultEventLoggerRepository(dbx *sqlx.DB) *EventLoggerRepository {
	eventLoggerRepository := &EventLoggerRepository{
		database: dbx,
	}
	return eventLoggerRepository
}

func (er *EventLoggerRepository) Add(event *event_logger_model.Event) error {

	// insert main request info
	result, err := er.database.Exec(`
		INSERT INTO gocms_plugin_event_logger_request (path, method, created) VALUES (?, ?, ?)
	`, event.Url, event.Method, event.Datetime)
	if err != nil {
		log.Errorf("Error adding event log to database: %s", err.Error())
		return err
	}
	// add id to user object
	id, _ := result.LastInsertId()
	event.Id = id

	// insert headers
	for header, values := range event.Headers {
		vals := fmt.Sprintf("%s", values)
		result, err = er.database.Exec(`
			INSERT INTO gocms_plugin_event_logger_request_headers (requestId, header, content, created) VALUES (?, ?, ?, ?)
		`, event.Id, header, vals, event.Datetime)
		if err != nil {
			log.Errorf("Error adding event log to database: %s", err.Error())
		}
	}
	return err // will be nil if no error

}
