package event_logger_model

import (
	"time"
)

type Event struct {
	Id       int64
	Url      string
	Method   string
	Headers  map[string][]string
	Datetime time.Time
}
