package entity

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/common/aggregate"
	commonevt "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/common/event"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity/valueobject"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/event"
)

// LeaveAggregate aggregate
type LeaveAggregate struct {
	*aggregate.AggregateBase
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

var _ aggregate.Aggregate = (*LeaveAggregate)(nil)

func NewLeaveAggregateWithID(id string) *LeaveAggregate {
	aggregate := NewLeaveAggregate()
	aggregate.SetID(id)
	return aggregate
}

func NewLeaveAggregate() *LeaveAggregate {
	leaveAggregate := &LeaveAggregate{}
	base := aggregate.NewAggregateBase()
	leaveAggregate.AggregateBase = base
	return leaveAggregate
}

func (agg *LeaveAggregate) On(evt commonevt.Event) error {
	switch evt.Type() {
	case event.LeaveCreated:
		return agg.OnLeaveCreated(evt)
	default:
		return fmt.Errorf("invalid event type: %s", evt.Type())
	}
}

func (agg *LeaveAggregate) GetDuration() int64 {
	duration := agg.EndTime.Sub(agg.StartTime)
	return int64(duration)
}

func (agg *LeaveAggregate) AddHistoryApprovalInfo(approvalInfo *ApprovalInfo) {
	if agg.HistoryApprovalInfos == nil {
		agg.HistoryApprovalInfos = make([]*ApprovalInfo, 0)
	}

	agg.HistoryApprovalInfos = append(agg.HistoryApprovalInfos, approvalInfo)
}

func (agg *LeaveAggregate) Create() {
	agg.StartTime = time.Now()
	agg.Status = valueobject.APPROVING
}

func (agg *LeaveAggregate) Agree(nextApprover valueobject.Approver) {
	agg.Approver = nextApprover
	agg.Status = valueobject.APPROVING
}

func (agg *LeaveAggregate) Reject(approver valueobject.Approver) {
	agg.Approver = approver
	agg.Status = valueobject.REJECTED
	agg.Approver = valueobject.Approver{}
}

func (agg *LeaveAggregate) Finish() {
	agg.Approver = valueobject.Approver{}
	agg.Status = valueobject.APPROVED
	agg.EndTime = time.Now()
	agg.Duration = int64(agg.EndTime.Sub(agg.StartTime))
}

func (agg *LeaveAggregate) OnLeaveCreated(evt commonevt.Event) error {
	var dat event.LeaveCreatedEvent
	if err := json.Unmarshal(evt.Data(), &dat); err != nil {
		return err
	}

	return nil
}
