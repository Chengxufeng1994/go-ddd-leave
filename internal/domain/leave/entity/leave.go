package entity

// Leave aggregate
type Leave struct {
	CurrentApprovalInfo *ApprovalInfo
}

func (leave *Leave) Create(approvalInfo *ApprovalInfo) Leave {
	return Leave{}
}
