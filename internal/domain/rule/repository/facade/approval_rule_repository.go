package facade

import "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/rule/entity"

type ApprovalRuleRepository interface {
	GetLeaderMaxLevel(rule *entity.ApprovalRule) int
}
