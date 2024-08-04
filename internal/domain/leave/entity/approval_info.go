package entity

import (
	"time"

	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity/valueobject"
)

// entity
type ApprovalInfo struct {
	ApprovalInfoID string
	Approver       valueobject.Approver
	ApprovalType   valueobject.ApprovalType
	Message        string
	Time           time.Time
}
