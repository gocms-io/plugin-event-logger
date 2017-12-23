package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gocms-io/plugin-event-logger/domain/health/health_controller"
	"github.com/gocms-io/plugin-event-logger/init/service"
	"github.com/gocms-io/plugin-event-logger/middleware/user_context"
	"github.com/gocms-io/plugin-event-logger/middleware/timezone"
	"github.com/gocms-io/plugin-event-logger/domain/event_logger/event_logger_controller"
)

type ControllersGroup struct {
	apiControllers     *ApiControllers
	contentControllers *ContentControllers
}

type ApiControllers struct {
	healthController          *health_controller.HealthController
	eventLoggerController *event_logger_controller.EventLoggerController
}

type ContentControllers struct {
	DocumentationController *DocumentationController
}

func DefaultControllerGroup(r *gin.Engine, servicesGroup *service.ServicesGroup) *ControllersGroup {

	// todo the timezone doesn't seem to pull through

	r.Use(timezoneMdl.Timezone())
	r.Use(user_middleware.UserHeaderContext())

	controllersGroup := &ControllersGroup{
		apiControllers: &ApiControllers{
			healthController:          health_controller.DefaultHealthController(r, servicesGroup),
			eventLoggerController: event_logger_controller.DefaultEventLoggerController(r, servicesGroup),
		},
		contentControllers: &ContentControllers{
			DocumentationController: DefaultDocumentationController(r),
		},
	}

	return controllersGroup
}
