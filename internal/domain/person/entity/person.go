package entity

import (
	"time"

	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/entity/valueobject"
)

type Person struct {
	PersonID         string
	PersonName       string
	PersonType       valueobject.PersonType
	RelationShip     []*RelationShip
	RoleLevel        int
	CreateTime       time.Time
	LastModifiedTime time.Time
	Status           valueobject.PersonStatus
}

func (entity *Person) Create() {
	entity.CreateTime = time.Now()
	entity.Status = valueobject.ENABLE
}

func (entity *Person) Enable() {
	entity.LastModifiedTime = time.Now()
	entity.Status = valueobject.ENABLE
}

func (entity *Person) Disable() {
	entity.LastModifiedTime = time.Now()
	entity.Status = valueobject.DISABLE
}
