package health_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gocms-io/plugin-event-logger/init/service"
	"github.com/gocms-io/plugin-event-logger/utility/errors"
)

type HealthController struct {
	Routes       *gin.RouterGroup
	serviceGroup *service.ServicesGroup
}

func DefaultHealthController(r *gin.Engine, serviceGroup *service.ServicesGroup) *HealthController {
	hc := &HealthController{
		Routes:       r.Group("/api"),
		serviceGroup: serviceGroup,
	}

	hc.Default()
	return hc
}

func (hc *HealthController) Default() {
	hc.Routes.GET("/healthy", hc.health)
}

/**
* @api {get} /healthy Service Health Status
* @apiDescription Used to verify that the services are up and running.
* @apiName GetHealthy
* @apiGroup Utility
 */
func (hc *HealthController) health(c *gin.Context) {

	ok, _ := hc.serviceGroup.HealthService.GetHealthStatus()

	if !ok {

		msg := "Service is having health issues"

		// todo add a admin key to get more detailed reports
		//for i, issue := range context {
		//	if i == 0 {
		//		msg = fmt.Sprintf("%v %v", msg, issue)
		//	} else {
		//		msg = fmt.Sprintf("%v, %v", msg, issue)
		//	}
		//}

		errors.Response(c, http.StatusInternalServerError, msg, nil)
		return
	}

	c.Status(http.StatusOK)
}
