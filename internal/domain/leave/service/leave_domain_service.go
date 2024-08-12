package service

import (
	"context"
	"fmt"

	commonevent "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/common/event"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity/valueobject"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/event"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/repository/facade"
)

type LeaveDomainService struct {
	leaveFactory        LeaveFactory
	leaveRepository     facade.LeaveRepositoryInterface
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
func (leaveDomainService *LeaveDomainService) CreateLeave(leave *entity.LeaveAggregate, leaderMaxLevel int, approver valueobject.Approver) error {
	leave.LeaderMaxLevel = leaderMaxLevel
	leave.Approver = approver
	leave.Create()
	leaveDomainService.leaveRepository.Save(context.Background(), leave)
	evt := event.NewLeaveCreatedEvent(leave)
	leaveDomainService.leaveRepository.SaveEvent(evt)
	leaveDomainService.leaveEventPublisher.Publish(evt)
	return nil
}

// TODO: trx
func (leaveDomainService *LeaveDomainService) UpdateLeave(leave *entity.LeaveAggregate) error {
	existed, err := leaveDomainService.leaveRepository.Load(context.Background(), leave.GetID())
	if err != nil {
		return fmt.Errorf("leave %s not found", leave.GetID())
	}
	if existed == nil {
		return fmt.Errorf("leave %s not found", leave.GetID())
	}

	leaveDomainService.leaveRepository.Save(context.Background(), leave)
	return nil
}

func (leaveDomainService *LeaveDomainService) SubmitApproval(leave *entity.LeaveAggregate, approver valueobject.Approver) error {
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
	leaveDomainService.leaveRepository.Save(context.Background(), leave)
	leaveDomainService.leaveRepository.SaveEvent(evt)
	leaveDomainService.leaveEventPublisher.Publish(evt)

	return nil
}

func (leaveDomainService *LeaveDomainService) GetLeaveInfo(leaveID string) (*entity.LeaveAggregate, error) {
	leave, err := leaveDomainService.leaveRepository.Load(context.Background(), leaveID)
	if err != nil {
		return nil, err
	}
	return leave, nil
}

func (leaveDomainService *LeaveDomainService) QueryLeaveInfoByApplicant(applicantID string) ([]*entity.LeaveAggregate, error) {
	leaveList := leaveDomainService.leaveRepository.QueryByApplicantID(context.Background(), applicantID)
	return leaveList, nil
}

func (leaveDomainService *LeaveDomainService) QueryLeaveInfoByApprover(approverID string) ([]*entity.LeaveAggregate, error) {
	leaveList := leaveDomainService.leaveRepository.QueryByApproverID(context.Background(), approverID)
	return leaveList, nil
}
