package po

import (
	"time"
)

type LeaveEvent struct {
	ID             string
	Source         string
	LeaveEventType string
	Data           []byte
	Timestamp      time.Time
}
