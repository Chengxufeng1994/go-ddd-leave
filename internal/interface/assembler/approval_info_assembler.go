package assembler

import (
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/interface/dto"
)

type ApprovalInfoAssembler struct {
	approverAssembler *ApproverAssembler
}

func NewApprovalInfoAssembler() *ApprovalInfoAssembler {
	return &ApprovalInfoAssembler{
		approverAssembler: NewApproverAssembler(),
	}
}

func (assembler *ApprovalInfoAssembler) ToDTO(do *entity.ApprovalInfo) dto.ApprovalInfoDTO {
	if do == nil {
		return dto.ApprovalInfoDTO{}
	}

	return dto.ApprovalInfoDTO{
		ApprovalInfoID: do.ApprovalInfoID,
		Message:        do.Message,
		Time:           do.Time,
	}
}

func (assembler *ApprovalInfoAssembler) ToDO(dto *dto.ApprovalInfoDTO) *entity.ApprovalInfo {
	approvalInfo := new(entity.ApprovalInfo)
	approvalInfo.ApprovalInfoID = dto.ApprovalInfoID
	approvalInfo.Approver = assembler.approverAssembler.ToDO(&dto.ApproverDTO)
	approvalInfo.Message = dto.Message
	return approvalInfo
}
