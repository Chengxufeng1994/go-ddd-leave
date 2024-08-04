package mapper

import (
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/rule/entity"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/rule/repository/po"
)

type ApprovalRuleMapper struct{}

func NewApprovalRuleMapper() ApprovalRuleMapper {
	return ApprovalRuleMapper{}
}

func (m ApprovalRuleMapper) ToDomain(po *po.ApprovalRule) *entity.ApprovalRule {
	return &entity.ApprovalRule{
		LeaveType:      po.LeaveType,
		PersonType:     po.PersonType,
		Duration:       po.Duration,
		MaxLeaderLevel: po.MaxLeaderLevel,
	}
}

func (m ApprovalRuleMapper) ToPersistence(do *entity.ApprovalRule) *po.ApprovalRule {
	return &po.ApprovalRule{
		LeaveType:      do.LeaveType,
		PersonType:     do.PersonType,
		Duration:       do.Duration,
		MaxLeaderLevel: do.MaxLeaderLevel,
	}
}
