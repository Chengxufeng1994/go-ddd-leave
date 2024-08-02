package facade

import "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/repository/po"

type PersonRepositoryInterface interface {
	Insert(person *po.Person)
	Update(person *po.Person)
	FindByID(id string) (*po.Person, error)
	FindByLeaderByPersonID(leaderID string) (*po.Person, error)
}
