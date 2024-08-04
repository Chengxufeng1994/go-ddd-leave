package assembler

import (
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/entity"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/entity/valueobject"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/interface/dto"
)

type PersonAssembler struct{}

func NewPersonAssembler() *PersonAssembler {
	return &PersonAssembler{}
}

func (assembler *PersonAssembler) ToDTO(do *entity.Person) *dto.PersonDTO {
	return &dto.PersonDTO{
		PersonID:         do.PersonID,
		PersonName:       do.PersonName,
		PersonType:       int(do.PersonType),
		RoleID:           do.RoleLevel,
		CreateTime:       do.CreateTime,
		LastModifiedTime: do.LastModifiedTime,
		Status:           int(do.Status),
	}
}

func (assembler *PersonAssembler) ToDO(dto *dto.PersonDTO) *entity.Person {
	return &entity.Person{
		PersonID:         dto.PersonID,
		PersonName:       dto.PersonName,
		RoleLevel:        dto.RoleID,
		PersonType:       valueobject.NewPersonType(dto.PersonType),
		CreateTime:       dto.CreateTime,
		LastModifiedTime: dto.LastModifiedTime,
		Status:           valueobject.NewPersonStatus(dto.Status),
	}
}
