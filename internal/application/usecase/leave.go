package usecase

import "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity"

type LeaveUseCase interface {
	CreateLeaveInfo(entity *entity.LeaveAggregate)
	GetLeaveInfo(leaveID string) (*entity.LeaveAggregate, error)
	UpdateLeaveInfo(entity *entity.LeaveAggregate)
	SubmitApproval(entity *entity.LeaveAggregate)
	QueryLeaveInfosByApplicant(applicantID string) ([]*entity.LeaveAggregate, error)
	QueryLeaveInfosByApprover(approverID string) ([]*entity.LeaveAggregate, error)
}
