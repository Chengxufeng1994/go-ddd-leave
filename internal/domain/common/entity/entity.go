package entity

type Entity interface {
	GetID() string
	SetID(string)
}

type EntityBase struct {
	id string
}

func NewEntityBase() *EntityBase {
	return &EntityBase{}
}

func (entity *EntityBase) GetID() string {
	return entity.id
}

func (entity *EntityBase) SetID(id string) {
	entity.id = id
}
