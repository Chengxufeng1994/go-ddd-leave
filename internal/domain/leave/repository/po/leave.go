package po

import "time"

type Leave struct {
	ID                        string
	ApplicantID               string
	ApplicantName             string
	ApplicantType             string
	ApproverID                string
	ApproverName              string
	LeaveType                 string
	Status                    string
	StartTime                 time.Time
	EndTime                   time.Time
	Duration                  int64
	HistoryApprovalInfoPOList []*ApprovalInfo
}
