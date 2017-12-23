package timezoneMdl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"github.com/gocms-io/gocms/context/consts"
)

const TIMEZONE_MIDDLEWARE_KEY = "TIMEZONE_MIDDLEWARE_KEY"

// Timezone setups the default timezone middleware
// to extract a timezone out of the header or default to Local
// utilize timezones as defined here: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
func Timezone() gin.HandlerFunc {
	return timezoneMiddleware
}

// GetTimezoneFromContext gets the timezone set in the request context.
func GetTimezoneFromContext(c *gin.Context) (*time.Location, bool) {
	// get timezone from context
	if timezone, ok := c.Get("TIMEZONE_MIDDLEWARE_KEY"); ok {
		if value, ok := timezone.(time.Location); ok {
			return &value, true
		}
	}
	// if error get local and return
	local, _ := time.LoadLocation("Local")
	return local, false
}

func timezoneMiddleware(c *gin.Context) {
	timezoneHeader := c.Request.Header.Get(consts.GOCMS_HEADER_TIMEZONE_KEY)
	if timezoneHeader == "" {
		timezoneHeader = "Local"
	}
	timezone, err := time.LoadLocation(timezoneHeader)
	if err != nil {
		fmt.Printf("Error parsing timezone header %v: %v\n", timezoneHeader, err)
		// todo currently default to est because we are in GR. Change in the future
		timezone, _ = time.LoadLocation("America/Detroit")
	}
	c.Set(TIMEZONE_MIDDLEWARE_KEY, *timezone)

	c.Next()
}
