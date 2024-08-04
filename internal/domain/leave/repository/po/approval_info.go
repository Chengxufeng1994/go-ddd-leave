package po

type ApprovalInfo struct {
	ApprovalInfoID string
	LeaveID        string
	ApplicantID    string
	ApproverID     string
	ApproverLevel  int
	ApproverName   string
	Message        string
	Time           int64
}
