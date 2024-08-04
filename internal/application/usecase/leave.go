package usecase

import "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity"

type LeaveUseCase interface {
	CreateLeaveInfo(entity *entity.Leave)
	GetLeaveInfo(leaveID string) (*entity.Leave, error)
	UpdateLeaveInfo(entity *entity.Leave)
	SubmitApproval(entity *entity.Leave)
	QueryLeaveInfosByApplicant(applicantID string) ([]*entity.Leave, error)
	QueryLeaveInfosByApprover(approverID string) ([]*entity.Leave, error)
}
