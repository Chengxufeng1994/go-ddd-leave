package service

import (
	"github.com/Chengxufeng1994/go-ddd-leave/internal/application/usecase"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity/valueobject"
	domainleaveservice "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/service"
	domainpersonservice "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/service"
	domainruleservice "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/rule/service"
)

type LeaveApplicationService struct {
	leaveDomainService        *domainleaveservice.LeaveDomainService
	personDomainService       *domainpersonservice.PersonDomainService
	approvalruleDomainService *domainruleservice.ApprovalRuleDomainService
}

func NewLeaveApplicationService(leaveDomainService *domainleaveservice.LeaveDomainService, personDomainService *domainpersonservice.PersonDomainService, approvalruleDomainService *domainruleservice.ApprovalRuleDomainService) usecase.LeaveUseCase {
	return &LeaveApplicationService{
		leaveDomainService:        leaveDomainService,
		personDomainService:       personDomainService,
		approvalruleDomainService: approvalruleDomainService,
	}
}

// CreateLeaveInfo implements usecase.LeaveUseCase.
func (svc *LeaveApplicationService) CreateLeaveInfo(leave *entity.LeaveAggregate) {
	leaderMaxLevel := svc.approvalruleDomainService.GetLeaderMaxLevel(leave.Applicant.PersonType, string(leave.LeaveType), leave.GetDuration())
	approver := svc.personDomainService.FindFirstApprover(leave.Applicant.PersonID, leaderMaxLevel)
	svc.leaveDomainService.CreateLeave(leave, leaderMaxLevel, valueobject.FromPerson(approver))
}

// UpdateLeaveInfo implements usecase.LeaveUseCase.
func (svc *LeaveApplicationService) UpdateLeaveInfo(leave *entity.LeaveAggregate) {
	svc.leaveDomainService.UpdateLeave(leave)
}

// GetLeaveInfo implements usecase.LeaveUseCase.
func (svc *LeaveApplicationService) GetLeaveInfo(leaveID string) (*entity.LeaveAggregate, error) {
	return svc.leaveDomainService.GetLeaveInfo(leaveID)
}

// SubmitApproval implements usecase.LeaveUseCase.
func (svc *LeaveApplicationService) SubmitApproval(leave *entity.LeaveAggregate) {
	approver := svc.personDomainService.FindNextApprover(leave.Approver.PersonID, leave.LeaderMaxLevel)
	svc.leaveDomainService.SubmitApproval(leave, valueobject.FromPerson(approver))
}

// QueryLeaveInfosByApplicant implements usecase.LeaveUseCase.
func (svc *LeaveApplicationService) QueryLeaveInfosByApplicant(applicantID string) ([]*entity.LeaveAggregate, error) {
	return svc.leaveDomainService.QueryLeaveInfoByApplicant(applicantID)
}

// QueryLeaveInfosByApprover implements usecase.LeaveUseCase.
func (svc *LeaveApplicationService) QueryLeaveInfosByApprover(approverID string) ([]*entity.LeaveAggregate, error) {
	return svc.leaveDomainService.QueryLeaveInfoByApprover(approverID)
}
