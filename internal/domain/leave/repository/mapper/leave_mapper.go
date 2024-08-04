package mapper

import (
	"time"

	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity/valueobject"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/repository/po"
)

type LeaveMapper struct{}

func NewLeaveMapper() LeaveMapper {
	return LeaveMapper{}
}

func (m LeaveMapper) ToDomain(po *po.Leave) *entity.Leave {
	do := &entity.Leave{
		Applicant:            valueobject.NewApplicant(po.ApplicantID, po.ApplicantName, po.ApplicantType),
		Approver:             valueobject.NewApprover(po.ApproverID, po.ApproverName),
		LeaveType:            valueobject.NewLeaveType(po.LeaveType),
		Status:               valueobject.NewStatus(po.Status),
		StartTime:            po.StartTime,
		EndTime:              po.EndTime,
		HistoryApprovalInfos: m.approvalInfoListDOFromPO(po),
		// EndTime:   po.EndTime,
		// Duration:  po.Duration,
	}

	do.SetID(po.ID)
	return do
}

func (m LeaveMapper) ToCreateLeavePersistence(do *entity.Leave) *po.Leave {
	leavePo := new(po.Leave)
	leavePo.ID = do.GetID()
	leavePo.ApplicantID = do.Applicant.PersonID
	leavePo.ApplicantName = do.Applicant.PersonName
	leavePo.ApproverID = do.Approver.PersonID
	leavePo.ApproverName = do.Approver.PersonName
	leavePo.StartTime = do.StartTime
	leavePo.Status = string(do.Status)
	leavePo.HistoryApprovalInfoPOList = m.approvalInfoListPOFromDO(do)
	return leavePo
}

func (m LeaveMapper) ToPersistence(do *entity.Leave) *po.Leave {
	return &po.Leave{
		ID:                        do.GetID(),
		ApplicantID:               do.Applicant.PersonID,
		ApplicantName:             do.Applicant.PersonName,
		ApplicantType:             do.Applicant.PersonType,
		ApproverID:                do.Approver.PersonID,
		ApproverName:              do.Approver.PersonName,
		LeaveType:                 string(do.LeaveType),
		StartTime:                 do.StartTime,
		EndTime:                   do.EndTime,
		Status:                    string(do.Status),
		HistoryApprovalInfoPOList: m.approvalInfoListPOFromDO(do),
		// EndTime:       do.EndTime,
		// Duration:      do.Duration,
	}
}

func (m LeaveMapper) approvalInfoListPOFromDO(do *entity.Leave) []*po.ApprovalInfo {
	n := len(do.HistoryApprovalInfos)
	historyApprovalInfos := make([]*po.ApprovalInfo, 0, n)
	for i := 0; i < n; i++ {
		historyApprovalInfos = append(historyApprovalInfos, m.approvalInfoPOFromDO(do.HistoryApprovalInfos[i]))
	}

	return historyApprovalInfos
}

func (m LeaveMapper) approvalInfoPOFromDO(do *entity.ApprovalInfo) *po.ApprovalInfo {
	return &po.ApprovalInfo{
		ApproverID:     do.Approver.PersonID,
		ApproverLevel:  do.Approver.Level,
		ApproverName:   do.Approver.PersonName,
		ApprovalInfoID: do.ApprovalInfoID,
		Message:        do.Message,
		Time:           do.Time.Unix(),
	}
}

func (m LeaveMapper) approvalInfoListDOFromPO(po *po.Leave) []*entity.ApprovalInfo {
	n := len(po.HistoryApprovalInfoPOList)
	historyApprovalInfos := make([]*entity.ApprovalInfo, 0, n)
	for i := 0; i < n; i++ {
		historyApprovalInfos = append(historyApprovalInfos, m.approvalInfoDOFromPO(po.HistoryApprovalInfoPOList[i]))
	}

	return historyApprovalInfos
}

func (m LeaveMapper) approvalInfoDOFromPO(po *po.ApprovalInfo) *entity.ApprovalInfo {
	return &entity.ApprovalInfo{
		ApprovalInfoID: po.ApprovalInfoID,
		Approver: valueobject.Approver{
			PersonID:   po.ApproverID,
			PersonName: po.ApproverName,
			Level:      po.ApproverLevel},
		Message: po.Message,
		Time:    time.Unix(po.Time, 0),
	}
}
