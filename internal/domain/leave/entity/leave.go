package entity

import (
	"time"

	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/common/aggregate"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity/valueobject"
)

// Leave aggregate
type Leave struct {
	aggregate.AggregateBase
	Applicant valueobject.Applicant
	Approver  valueobject.Approver
	LeaveType valueobject.LeaveType
	Status    valueobject.Status
	StartTime time.Time
	EndTime   time.Time
	Duration  int64

	LeaderMaxLevel       int // 批准的最高等級
	CurrentApprovalInfo  *ApprovalInfo
	HistoryApprovalInfos []*ApprovalInfo
}

var _ aggregate.Aggregate = (*Leave)(nil)

func (entity *Leave) GetDuration() int64 {
	duration := entity.EndTime.Sub(entity.StartTime)
	return int64(duration)
}

func (entity *Leave) AddHistoryApprovalInfo(approvalInfo *ApprovalInfo) {
	if entity.HistoryApprovalInfos == nil {
		entity.HistoryApprovalInfos = make([]*ApprovalInfo, 0)
	}

	entity.HistoryApprovalInfos = append(entity.HistoryApprovalInfos, approvalInfo)
}

func (entity *Leave) Create() {
	entity.StartTime = time.Now()
	entity.Status = valueobject.APPROVING
}

func (entity *Leave) Agree(nextApprover valueobject.Approver) {
	entity.Approver = nextApprover
	entity.Status = valueobject.APPROVING
}

func (entity *Leave) Reject(approver valueobject.Approver) {
	entity.Approver = approver
	entity.Status = valueobject.REJECTED
	entity.Approver = valueobject.Approver{}
}

func (entity *Leave) Finish() {
	entity.Approver = valueobject.Approver{}
	entity.Status = valueobject.APPROVED
	entity.EndTime = time.Now()
	entity.Duration = int64(entity.EndTime.Sub(entity.StartTime))
}
