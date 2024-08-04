package po

type ApprovalRule struct {
	ID             string
	LeaveType      string
	PersonType     string
	Duration       int64
	MaxLeaderLevel int

	ApplicantRoleID string
}
