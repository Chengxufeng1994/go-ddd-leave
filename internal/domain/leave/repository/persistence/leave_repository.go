package persistence

import "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/repository/facade"

type LeaveRepository struct{}

var _ facade.LeaveRepositoryInterface = (*LeaveRepository)(nil)

func NewLeaveRepository() facade.LeaveRepositoryInterface {
	return &LeaveRepository{}
}
