package entity

import (
	"time"

	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/valueobject"
)

type Person struct {
	PersonID         string
	PersonName       string
	PersonType       valueobject.PersonType
	RelationShip     RelationShip
	RoleLevel        int
	createTime       time.Time
	LastModifiedTime time.Time
}

func (p *Person) Create() {}
