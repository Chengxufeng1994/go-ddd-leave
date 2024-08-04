package dao

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/repository/po"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ApprovalInfoDao interface {
	SaveAll(list []*po.ApprovalInfo)
	QueryByLeaveID(leaveID string) []*po.ApprovalInfo
}

type ApprovalInfoDaoImpl struct {
	db *gorm.DB
}

var _ ApprovalInfoDao = (*ApprovalInfoDaoImpl)(nil)

func NewApprovalInfoDao(db *gorm.DB) ApprovalInfoDao {
	return &ApprovalInfoDaoImpl{
		db: db,
	}
}

// SaveAll implements ApprovalInfoDao.
func (dao *ApprovalInfoDaoImpl) SaveAll(list []*po.ApprovalInfo) {
	for _, item := range list {
		dao.db.WithContext(context.Background()).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Model(&po.ApprovalInfo{}).
			Save(&item)
	}
}

// QueryByLeaveID implements ApprovalInfoDao.
func (dao *ApprovalInfoDaoImpl) QueryByLeaveID(leaveID string) []*po.ApprovalInfo {
	list := make([]*po.ApprovalInfo, 0)
	dao.db.WithContext(context.Background()).
		Model(&po.ApprovalInfo{}).
		Where("leave_id = ?", leaveID).
		Find(&list)

	return list
}
