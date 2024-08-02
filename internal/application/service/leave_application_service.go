package service

import "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/service"

type LeaveApplicationService struct {
	leaveDomainService *service.LeaveDomainService
}
