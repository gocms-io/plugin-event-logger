package api_utility

import (
	"net/http"
	"io"
	"github.com/gin-gonic/gin"
	"github.com/gocms-io/gocms/utility/log"
	"github.com/gocms-io/plugin-event-logger/utility/errors"
)

func SendCarbonCopyRequest(c *gin.Context) {
	// get request headers and add them to response
	resHeaders := c.Writer.Header()
	for header, values := range c.Request.Header {
		resHeaders[header] = values
	}
	c.Status(http.StatusOK)
	_, err := io.Copy(c.Writer, c.Request.Body)
	// error with copying body
	if err != nil {
		log.Errorf("Error Copying Body: %v\n", err.Error())
		errors.Response(c, http.StatusBadGateway, "Error Copying Body", err.Error())
		return
	}
	c.Done()
}