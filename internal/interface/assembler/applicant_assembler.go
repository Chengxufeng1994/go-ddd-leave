package assembler

import (
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/entity/valueobject"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/interface/dto"
)

type ApplicantAssembler struct{}

func NewApplicantAssembler() *ApplicantAssembler {
	return &ApplicantAssembler{}
}

func (assembler *ApplicantAssembler) ToDTO(do *valueobject.Applicant) dto.ApplicantDTO {
	return dto.ApplicantDTO{
		PersonID:   do.PersonID,
		PersonName: do.PersonName,
	}
}

func (assembler *ApplicantAssembler) ToDO(dto *dto.ApplicantDTO) valueobject.Applicant {
	return valueobject.Applicant{
		PersonID:   dto.PersonID,
		PersonName: dto.PersonName,
	}
}
