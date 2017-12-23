package event_logger_model

import (
	"time"
)

type Event struct {
	Url      string
	Headers  map[string]string
	Datetime time.Time
}
