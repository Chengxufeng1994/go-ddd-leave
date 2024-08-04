package service

import (
	"fmt"

	commonevent "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/common/event"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity/valueobject"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/event"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/repository/facade"
)

type LeaveDomainService struct {
	leaveRepository     facade.LeaveRepositoryInterface
	leaveFactory        LeaveFactory
	leaveEventPublisher commonevent.EventPublisher
}

func NewLeaveDomainService(leaveRepository facade.LeaveRepositoryInterface, leaveFactory LeaveFactory, leaveEventPublisher commonevent.EventPublisher) *LeaveDomainService {
	return &LeaveDomainService{
		leaveRepository:     leaveRepository,
		leaveFactory:        leaveFactory,
		leaveEventPublisher: leaveEventPublisher,
	}
}

// TODO: trx
func (leaveDomainService *LeaveDomainService) CreateLeave(leave *entity.Leave, leaderMaxLevel int, approver valueobject.Approver) error {
	leave.LeaderMaxLevel = leaderMaxLevel
	leave.Approver = approver
	leave.Create()
	leaveDomainService.leaveRepository.Save(leave)
	evt := event.NewLeaveCreatedEvent(leave)
	leaveDomainService.leaveRepository.SaveEvent(evt)
	leaveDomainService.leaveEventPublisher.Publish(evt)
	return nil
}

// TODO: trx
func (leaveDomainService *LeaveDomainService) UpdateLeave(leave *entity.Leave) error {
	existed, err := leaveDomainService.leaveRepository.FindByID(leave.GetID())
	if err != nil {
		return fmt.Errorf("leave %s not found", leave.GetID())
	}
	if existed == nil {
		return fmt.Errorf("leave %s not found", leave.GetID())
	}

	leaveDomainService.leaveRepository.Save(leave)
	return nil
}

func (leaveDomainService *LeaveDomainService) SubmitApproval(leave *entity.Leave, approver valueobject.Approver) error {
	var evt event.LeaveEvent
	if valueobject.REJECT == leave.CurrentApprovalInfo.ApprovalType {
		leave.Reject(approver)
		evt = event.NewLeaveRejectedEvent(leave)
	} else {
		if !approver.IsEmpty() {
			leave.Agree(approver)
		} else {
			leave.Finish()
		}
	}

	leave.AddHistoryApprovalInfo(leave.CurrentApprovalInfo)
	leaveDomainService.leaveRepository.Save(leave)
	leaveDomainService.leaveRepository.SaveEvent(evt)
	leaveDomainService.leaveEventPublisher.Publish(evt)

	return nil
}

func (leaveDomainService *LeaveDomainService) GetLeaveInfo(leaveID string) (*entity.Leave, error) {
	leave, err := leaveDomainService.leaveRepository.FindByID(leaveID)
	if err != nil {
		return nil, err
	}
	return leave, nil
}

func (leaveDomainService *LeaveDomainService) QueryLeaveInfoByApplicant(applicantID string) ([]*entity.Leave, error) {
	leaveList := leaveDomainService.leaveRepository.QueryByApplicantID(applicantID)
	return leaveList, nil
}

func (leaveDomainService *LeaveDomainService) QueryLeaveInfoByApprover(approverID string) ([]*entity.Leave, error) {
	leaveList := leaveDomainService.leaveRepository.QueryByApproverID(approverID)
	return leaveList, nil
}
