package service

import (
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/rule/entity"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/rule/repository/facade"
)

type ApprovalRuleDomainService struct {
	approvalRuleRepository facade.ApprovalRuleRepository
}

func NewApprovalRuleDomainService(approvalRuleRepository facade.ApprovalRuleRepository) *ApprovalRuleDomainService {
	return &ApprovalRuleDomainService{
		approvalRuleRepository: approvalRuleRepository,
	}
}

func (svc *ApprovalRuleDomainService) GetLeaderMaxLevel(personType, leaveType string, duration int64) int {
	rule := entity.NewApprovalRule(personType, leaveType, duration)
	return svc.approvalRuleRepository.GetLeaderMaxLevel(rule)
}
