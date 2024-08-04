package persistence

import (
	"fmt"
	"strconv"

	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/event"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/repository/dao"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/repository/facade"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/repository/mapper"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/repository/po"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/infrastructure/util"
)

type LeaveRepository struct {
	leaveMapper     mapper.LeaveMapper
	leaveDao        dao.LeaveDao
	leaveEventDao   dao.LeaveEventDao
	approvalInfoDao dao.ApprovalInfoDao
	st              util.IDGenerator
}

var _ facade.LeaveRepositoryInterface = (*LeaveRepository)(nil)

func NewLeaveRepository(leaveDao dao.LeaveDao, approvalInfoDao dao.ApprovalInfoDao, leaveEventDao dao.LeaveEventDao, st util.IDGenerator) facade.LeaveRepositoryInterface {

	return &LeaveRepository{
		leaveMapper:     mapper.NewLeaveMapper(),
		leaveDao:        leaveDao,
		leaveEventDao:   leaveEventDao,
		approvalInfoDao: approvalInfoDao,
		st:              st,
	}
}

// Save implements facade.LeaveRepositoryInterface.
func (repo *LeaveRepository) Save(entity *entity.Leave) {
	stId, _ := repo.st.NextID()
	entity.SetID(strconv.Itoa(int(stId)))
	repo.leaveDao.Save(repo.leaveMapper.ToPersistence(entity))
	po := repo.leaveMapper.ToPersistence(entity)
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

// FindByID implements facade.LeaveRepositoryInterface.
func (repo *LeaveRepository) FindByID(id string) (*entity.Leave, error) {
	po := repo.leaveDao.FindByID(id)
	if po == nil {
		return nil, fmt.Errorf("leave %s not found", id)
	}
	return repo.leaveMapper.ToDomain(po), nil
}

// QueryByApplicantID implements facade.LeaveRepositoryInterface.
func (repo *LeaveRepository) QueryByApplicantID(applicantID string) []*entity.Leave {
	leaveList := repo.leaveDao.QueryByApplicantID(applicantID)
	for i := 0; i < len(leaveList); i++ {
		item := leaveList[i]
		approvalInfoList := repo.approvalInfoDao.QueryByLeaveID(item.ID)
		item.HistoryApprovalInfoPOList = approvalInfoList
	}

	entities := make([]*entity.Leave, 0, len(leaveList))
	for i := 0; i < len(leaveList); i++ {
		entities = append(entities, repo.leaveMapper.ToDomain(leaveList[i]))
	}
	return entities
}

// QueryByApproverID implements facade.LeaveRepositoryInterface.
func (repo *LeaveRepository) QueryByApproverID(approverID string) []*entity.Leave {
	leaveList := repo.leaveDao.QueryByApproverID(approverID)
	for i := 0; i < len(leaveList); i++ {
		item := leaveList[i]
		approvalInfoList := repo.approvalInfoDao.QueryByLeaveID(item.ID)
		item.HistoryApprovalInfoPOList = approvalInfoList
	}

	entities := make([]*entity.Leave, 0, len(leaveList))
	for i := 0; i < len(leaveList); i++ {
		entities = append(entities, repo.leaveMapper.ToDomain(leaveList[i]))
	}
	return entities
}
