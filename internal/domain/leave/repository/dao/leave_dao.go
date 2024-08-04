package dao

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/repository/po"
	"gorm.io/gorm"
)

type LeaveDao interface {
	Save(*po.Leave)
	FindByID(id string) *po.Leave
	QueryByApplicantID(applicantID string) []*po.Leave
	QueryByApproverID(approverID string) []*po.Leave
}

type GORMLeaveDao struct {
	db *gorm.DB
}

func NewLeaveDao(db *gorm.DB) LeaveDao {
	return &GORMLeaveDao{
		db: db,
	}
}

func (dao *GORMLeaveDao) FindByID(id string) *po.Leave {
	var leave po.Leave
	if err := dao.db.WithContext(context.Background()).
		Preload("HistoryApprovalInfoPOList").
		Model(&po.Leave{}).
		Where("id = ?", id).
		First(&leave).Error; err != nil {
		return nil
	}

	return &leave
}

// Save implements LeaveDao.
func (dao *GORMLeaveDao) Save(leave *po.Leave) {
	dao.db.WithContext(context.Background()).
		Model(&po.Leave{}).
		Where("id = ?", leave.ID).
		Save(&leave)
}

// QueryByApplicantID implements LeaveDao.
func (dao *GORMLeaveDao) QueryByApplicantID(applicantID string) []*po.Leave {
	var leaves []*po.Leave
	dao.db.WithContext(context.Background()).
		Model(&po.Leave{}).
		Where("applicant_id = ?", applicantID).
		Find(&leaves)

	return leaves
}

// QueryByApproverID implements LeaveDao.
func (dao *GORMLeaveDao) QueryByApproverID(approverID string) []*po.Leave {
	panic("unimplemented")
}
