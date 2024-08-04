package dto

import "time"

type LeaveDTO struct {
	LeaveID                    string          `json:"leave_id"`
	ApplicantDTO               ApplicantDTO    `json:"applicant"`
	ApproverDTO                ApproverDTO     `json:"approver"`
	LeaveType                  string          `json:"leave_type"`
	CurrentApprovalInfoDTO     ApprovalInfoDTO `json:"current_approval_info"`
	HistoryApprovalInfoDTOList []ApprovalInfoDTO
	StartTime                  time.Time
	EndTime                    time.Time
	Duration                   int64
	Status                     string
}
