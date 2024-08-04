package facade

import (
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/entity"
)

type PersonRepositoryInterface interface {
	Save(person *entity.Person)
	Update(person *entity.Person)
	FindByID(personID string) *entity.Person
	FindByLeaderByPersonID(leaderID string) *entity.Person
}
