package assembler

import (
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity/valueobject"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/interface/dto"
)

type LeaveAssembler struct {
	applicantAssembler    *ApplicantAssembler
	approverAssembler     *ApproverAssembler
	approvalInfoAssembler *ApprovalInfoAssembler
}

func NewLeaveAssembler() *LeaveAssembler {
	return &LeaveAssembler{
		applicantAssembler:    NewApplicantAssembler(),
		approverAssembler:     NewApproverAssembler(),
		approvalInfoAssembler: NewApprovalInfoAssembler(),
	}
}

// domain object to data transfer object
func (leaveAssembler *LeaveAssembler) ToDTO(do *entity.LeaveAggregate) *dto.LeaveDTO {
	historyApprovalInfoDTOList := make([]dto.ApprovalInfoDTO, 0, len(do.HistoryApprovalInfos))
	for i := 0; i < len(do.HistoryApprovalInfos); i++ {
		historyApprovalInfoDTOList = append(historyApprovalInfoDTOList, leaveAssembler.approvalInfoAssembler.ToDTO(do.HistoryApprovalInfos[i]))
	}

	return &dto.LeaveDTO{
		LeaveID:                    do.GetID(),
		LeaveType:                  string(do.LeaveType),
		Status:                     string(do.Status),
		StartTime:                  do.StartTime,
		EndTime:                    do.EndTime,
		Duration:                   do.Duration,
		ApplicantDTO:               leaveAssembler.applicantAssembler.ToDTO(&do.Applicant),
		ApproverDTO:                leaveAssembler.approverAssembler.ToDTO(&do.Approver),
		CurrentApprovalInfoDTO:     leaveAssembler.approvalInfoAssembler.ToDTO(do.CurrentApprovalInfo),
		HistoryApprovalInfoDTOList: historyApprovalInfoDTOList,
	}
}

// data transfer object to domain object
func (leaveAssembler *LeaveAssembler) ToDO(dto *dto.LeaveDTO) *entity.LeaveAggregate {
	leave := entity.NewLeaveAggregate()
	leave.SetID(dto.LeaveID)
	leave.LeaveType = valueobject.NewLeaveType(dto.LeaveType)
	leave.Applicant = leaveAssembler.applicantAssembler.ToDO(&dto.ApplicantDTO)
	leave.Approver = leaveAssembler.approverAssembler.ToDO(&dto.ApproverDTO)
	leave.CurrentApprovalInfo = leaveAssembler.approvalInfoAssembler.ToDO(&dto.CurrentApprovalInfoDTO)
	historyApprovalInfos := make([]*entity.ApprovalInfo, 0, len(dto.HistoryApprovalInfoDTOList))
	for i := 0; i < len(dto.HistoryApprovalInfoDTOList); i++ {
		historyApprovalInfos = append(historyApprovalInfos, leaveAssembler.approvalInfoAssembler.ToDO(&dto.HistoryApprovalInfoDTOList[i]))
	}
	leave.HistoryApprovalInfos = historyApprovalInfos
	return leave
}
