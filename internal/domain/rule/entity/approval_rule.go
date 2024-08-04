package entity

import "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity"

type ApprovalRule struct {
	PersonType     string
	LeaveType      string
	Duration       int64
	MaxLeaderLevel int
}

func NewApprovalRule(personType, leaveType string, duration int64) *ApprovalRule {
	approvalRule := new(ApprovalRule)
	approvalRule.PersonType = personType
	approvalRule.LeaveType = leaveType
	approvalRule.Duration = duration
	return approvalRule
}

func NewApprovalRuleByLeave(leave *entity.Leave) *ApprovalRule {
	approvalRule := new(ApprovalRule)
	approvalRule.PersonType = leave.Applicant.PersonType
	approvalRule.LeaveType = string(leave.LeaveType)
	approvalRule.Duration = leave.GetDuration()
	return approvalRule
}
