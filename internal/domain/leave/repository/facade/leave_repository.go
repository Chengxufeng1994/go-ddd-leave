package facade

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/event"
)

type LeaveRepositoryInterface interface {
	Load(ctx context.Context, id string) (*entity.LeaveAggregate, error)
	Save(ctx context.Context, entity *entity.LeaveAggregate)
	QueryByApplicantID(ctx context.Context, applicantID string) []*entity.LeaveAggregate
	QueryByApproverID(ctx context.Context, approverID string) []*entity.LeaveAggregate

	SaveEvent(event event.LeaveEvent)
}
