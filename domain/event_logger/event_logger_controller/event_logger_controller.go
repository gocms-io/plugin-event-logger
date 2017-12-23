package event_logger_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gocms-io/plugin-event-logger/init/service"
	"github.com/gocms-io/plugin-event-logger/utility/api_utility"
	"github.com/gocms-io/plugin-event-logger/domain/event_logger/event_logger_model"
	"time"
)

type EventLoggerController struct {
	MiddlewareReceiverRoot *gin.RouterGroup
	serviceGroup           *service.ServicesGroup
}

// DefaultEventLoggerController create the default event logger route
func DefaultEventLoggerController(r *gin.Engine, serviceGroup *service.ServicesGroup) *EventLoggerController {

	el := &EventLoggerController{
		MiddlewareReceiverRoot: r.Group("/middleware"),
		serviceGroup:           serviceGroup,
	}

	el.Default()
	return el
}

// serve default routes
func (elc *EventLoggerController) Default() {
	elc.MiddlewareReceiverRoot.Any("/:rank/*path", elc.any)
}

// serve event logger route
func (elc *EventLoggerController) any(c *gin.Context) {
	// send of ok before we actually log anything
	// failing to log should not stop further execution
	api_utility.SendCarbonCopyRequest(c)

	path := c.Request.RequestURI
	headers := c.Request.Header

	event := event_logger_model.Event{
		Headers: headers,
		Url: path,
		Method: c.Request.Method,
		Datetime: time.Now(),
	}

	elc.serviceGroup.EventLoggerService.AddEventToLog(&event)
	return
}
