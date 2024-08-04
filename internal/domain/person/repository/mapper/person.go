package mapper

import (
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/entity"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/entity/valueobject"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/repository/po"
)

type PersonMapper struct {
}

// Domain object map to Persistent object
func NewPersonMapper() PersonMapper {
	return PersonMapper{}
}

// to domain object
func (m PersonMapper) ToDomain(po *po.Person) *entity.Person {
	return &entity.Person{
		PersonID:         po.PersonID,
		PersonName:       po.PersonName,
		PersonType:       valueobject.NewPersonType(po.PersonType),
		RoleLevel:        po.RoleLevel,
		Status:           valueobject.NewPersonStatus(po.Status),
		CreateTime:       po.CreateTime,
		LastModifiedTime: po.LastModifiedTime,
	}
}

// to persistence object
func (m PersonMapper) ToPersistence(do *entity.Person) *po.Person {
	return &po.Person{
		PersonID:         do.PersonID,
		PersonName:       do.PersonName,
		PersonType:       int(do.PersonType),
		RoleLevel:        do.RoleLevel,
		Status:           int(do.Status),
		CreateTime:       do.CreateTime,
		LastModifiedTime: do.LastModifiedTime,
	}
}
