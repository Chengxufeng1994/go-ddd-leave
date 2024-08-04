package persistence

import (
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/rule/entity"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/rule/repository/dao"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/rule/repository/facade"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/rule/repository/mapper"
)

type ApprovalRuleRepository struct {
	approvalRuleDao    dao.ApprovalRuleDao
	approvalRuleMapper mapper.ApprovalRuleMapper
}

func NewApprovalRuleRepository(approvalRuleDao dao.ApprovalRuleDao) facade.ApprovalRuleRepository {
	return &ApprovalRuleRepository{
		approvalRuleDao: approvalRuleDao,
	}
}

// GetLeaderMaxLevel implements facade.ApprovalRuleRepository.
func (repo *ApprovalRuleRepository) GetLeaderMaxLevel(rule *entity.ApprovalRule) int {
	po := repo.approvalRuleDao.FindRule(rule.PersonType, rule.LeaveType, rule.Duration)
	return po.MaxLeaderLevel
}
