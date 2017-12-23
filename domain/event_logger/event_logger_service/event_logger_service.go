package event_logger_service

import (
	"github.com/gocms-io/plugin-event-logger/domain/event_logger/event_logger_model"
	"github.com/gocms-io/plugin-event-logger/init/repository"
	"github.com/gocms-io/plugin-event-logger/context"
	"regexp"
)

type IEventLoggerService interface {
	AddEventToLog(event *event_logger_model.Event) error
}

type EventLoggerService struct {
	repositoryGroup *repository.RepositoriesGroup
}

var middlewareUrlRegex = regexp.MustCompile(`^/middleware/\d{1,4}`)

func DefaultEventLoggerService(rg *repository.RepositoriesGroup) *EventLoggerService {

	return &EventLoggerService{
		repositoryGroup: rg,
	}

}

func (els *EventLoggerService) AddEventToLog(event *event_logger_model.Event) error {

	// remove headers we don't want to save
	for _, headerToSkip := range context.Config.DbVars.IgnoreHeaders {
		delete(event.Headers, headerToSkip)
	}


	// remove /middleware/:rank
	event.Url = middlewareUrlRegex.ReplaceAllString(event.Url, "")

	// send to database
	err := els.repositoryGroup.EventLoggerRepository.Add(event)

	return err
}

