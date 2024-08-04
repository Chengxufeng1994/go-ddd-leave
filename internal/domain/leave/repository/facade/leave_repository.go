package facade

import (
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/event"
)

type LeaveRepositoryInterface interface {
	Save(entity *entity.Leave)
	FindByID(id string) (*entity.Leave, error)
	QueryByApplicantID(applicantID string) []*entity.Leave
	QueryByApproverID(approverID string) []*entity.Leave

	SaveEvent(event event.LeaveEvent)
}
