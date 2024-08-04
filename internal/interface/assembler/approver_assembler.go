package assembler

import (
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity/valueobject"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/interface/dto"
)

type ApproverAssembler struct{}

func NewApproverAssembler() *ApproverAssembler {
	return &ApproverAssembler{}
}

func (assembler *ApproverAssembler) ToDTO(do *valueobject.Approver) dto.ApproverDTO {
	return dto.ApproverDTO{
		PersonID:   do.PersonID,
		PersonName: do.PersonName,
	}
}

func (assembler *ApproverAssembler) ToDO(dto *dto.ApproverDTO) valueobject.Approver {
	return valueobject.Approver{
		PersonID:   dto.PersonID,
		PersonName: dto.PersonName,
	}
}
