package persistence

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/common/aggregate"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/event"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/repository/dao"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/repository/facade"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/repository/mapper"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/repository/po"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/infrastructure/util"
)

type LeaveRepository struct {
	st              util.IDGenerator
	mapper          mapper.LeaveMapper
	leaveDao        dao.LeaveDao
	leaveEventDao   dao.LeaveEventDao
	approvalInfoDao dao.ApprovalInfoDao
	esStore         aggregate.EventStore
}

var _ facade.LeaveRepositoryInterface = (*LeaveRepository)(nil)

func NewLeaveRepository(leaveDao dao.LeaveDao, approvalInfoDao dao.ApprovalInfoDao, leaveEventDao dao.LeaveEventDao, st util.IDGenerator) facade.LeaveRepositoryInterface {
	return &LeaveRepository{
		st:              st,
		mapper:          mapper.NewLeaveMapper(),
		leaveDao:        leaveDao,
		leaveEventDao:   leaveEventDao,
		approvalInfoDao: approvalInfoDao,
	}
}

// Save implements facade.LeaveRepositoryInterface.
func (repo *LeaveRepository) Save(ctx context.Context, entity *entity.LeaveAggregate) {
	stId, _ := repo.st.NextID()
	entity.SetID(strconv.Itoa(int(stId)))
	repo.leaveDao.Save(repo.mapper.ToPersistence(entity))
	po := repo.mapper.ToPersistence(entity)
	repo.approvalInfoDao.SaveAll(po.HistoryApprovalInfoPOList)
}

// Save implements facade.LeaveRepositoryInterface.
func (repo *LeaveRepository) SaveEvent(entity event.LeaveEvent) {
	repo.leaveEventDao.Save(&po.LeaveEvent{
		ID:             entity.ID(),
		Source:         string(entity.Name()),
		LeaveEventType: string(entity.Type()),
		Data:           entity.Data(),
		Timestamp:      entity.Time(),
	})
}

// Load implements facade.LeaveRepositoryInterface.
func (repo *LeaveRepository) Load(ctx context.Context, id string) (*entity.LeaveAggregate, error) {
	po := repo.leaveDao.FindByID(id)
	if po == nil {
		return nil, fmt.Errorf("leave %s not found", id)
	}
	return repo.mapper.ToDomain(po), nil
}

// QueryByApplicantID implements facade.LeaveRepositoryInterface.
func (repo *LeaveRepository) QueryByApplicantID(ctx context.Context, applicantID string) []*entity.LeaveAggregate {
	leaveList := repo.leaveDao.QueryByApplicantID(applicantID)
	for i := 0; i < len(leaveList); i++ {
		item := leaveList[i]
		approvalInfoList := repo.approvalInfoDao.QueryByLeaveID(item.ID)
		item.HistoryApprovalInfoPOList = approvalInfoList
	}

	entities := make([]*entity.LeaveAggregate, 0, len(leaveList))
	for i := 0; i < len(leaveList); i++ {
		entities = append(entities, repo.mapper.ToDomain(leaveList[i]))
	}
	return entities
}

// QueryByApproverID implements facade.LeaveRepositoryInterface.
func (repo *LeaveRepository) QueryByApproverID(ctx context.Context, approverID string) []*entity.LeaveAggregate {
	leaveList := repo.leaveDao.QueryByApproverID(approverID)
	for i := 0; i < len(leaveList); i++ {
		item := leaveList[i]
		approvalInfoList := repo.approvalInfoDao.QueryByLeaveID(item.ID)
		item.HistoryApprovalInfoPOList = approvalInfoList
	}

	entities := make([]*entity.LeaveAggregate, 0, len(leaveList))
	for i := 0; i < len(leaveList); i++ {
		entities = append(entities, repo.mapper.ToDomain(leaveList[i]))
	}
	return entities
}
