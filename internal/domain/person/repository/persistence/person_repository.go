package persistence

import (
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/repository/facade"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/repository/po"
)

type PersonRepository struct{}

func NewPersonRepository() facade.PersonRepositoryInterface {
	return &PersonRepository{}
}

// FindByID implements facade.PersonRepositoryInterface.
func (p *PersonRepository) FindByID(id string) (*po.Person, error) {
	panic("unimplemented")
}

// FindByLeaderByPersonID implements facade.PersonRepositoryInterface.
func (p *PersonRepository) FindByLeaderByPersonID(leaderID string) (*po.Person, error) {
	panic("unimplemented")
}

// Insert implements facade.PersonRepositoryInterface.
func (p *PersonRepository) Insert(person *po.Person) {
	panic("unimplemented")
}

// Update implements facade.PersonRepositoryInterface.
func (p *PersonRepository) Update(person *po.Person) {
	panic("unimplemented")
}
