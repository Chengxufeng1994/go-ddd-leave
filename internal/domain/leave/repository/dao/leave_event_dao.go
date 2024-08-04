package dao

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/repository/po"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LeaveEventDao interface {
	Save(*po.LeaveEvent)
}

type GORMLeaveEventDao struct {
	db *gorm.DB
}

func NewLeaveEventDao(db *gorm.DB) LeaveEventDao {
	return &GORMLeaveEventDao{
		db: db,
	}
}

func (dao *GORMLeaveEventDao) Save(leaveEvent *po.LeaveEvent) {
	dao.db.WithContext(context.Background()).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Model(&po.LeaveEvent{}).
		Create(leaveEvent)
}
