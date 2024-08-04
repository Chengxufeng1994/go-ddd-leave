package dao

import (
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/rule/repository/po"
	"gorm.io/gorm"
)

type ApprovalRuleDao interface {
	FindRule(applicantRoleId, leaveType string, duration int64) *po.ApprovalRule
}

type ApprovalRuleDaoImpl struct {
	db *gorm.DB
}

var _ ApprovalRuleDao = (*ApprovalRuleDaoImpl)(nil)

func NewApprovalRuleDao(db *gorm.DB) ApprovalRuleDao {
	return &ApprovalRuleDaoImpl{
		db: db,
	}
}

// FindRule implements ApprovalRuleDao.
func (a *ApprovalRuleDaoImpl) FindRule(applicantRoleID string, leaveType string, duration int64) *po.ApprovalRule {
	var rule po.ApprovalRule
	a.db.Model(&po.ApprovalRule{}).
		Where("person_type = ? AND leave_type = ? AND duration = ?", "INTERNAL", leaveType, duration).
		First(&rule)
	return &rule
}
