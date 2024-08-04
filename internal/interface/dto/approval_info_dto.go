package dto

import "time"

type ApprovalInfoDTO struct {
	ApprovalInfoID string      `json:"approval_info_id"`
	ApproverDTO    ApproverDTO `json:"approver"`
	Message        string
	Time           time.Time
}
