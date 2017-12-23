package event_logger_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gocms-io/plugin-event-logger/init/service"
	"net/http"
	"github.com/gocms-io/gocms/utility/log"
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

	log.Debugf("Logged Event")
	c.Status(http.StatusOK)
	return
}
