package persistence

import (
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/entity"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/repository/dao"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/repository/facade"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/repository/mapper"
)

type PersonRepository struct {
	personDao    dao.PersonDao
	personMapper mapper.PersonMapper
}

func NewPersonRepository(personDao dao.PersonDao) facade.PersonRepositoryInterface {
	return &PersonRepository{
		personDao:    personDao,
		personMapper: mapper.NewPersonMapper(),
	}
}

// Insert implements facade.PersonRepositoryInterface.
func (p *PersonRepository) Save(person *entity.Person) {
	po := p.personMapper.ToPersistence(person)
	p.personDao.Save(po)
}

// Update implements facade.PersonRepositoryInterface.
func (p *PersonRepository) Update(person *entity.Person) {
	po := p.personMapper.ToPersistence(person)
	p.personDao.Save(po)
}

// FindByID implements facade.PersonRepositoryInterface.
func (p *PersonRepository) FindByID(personID string) *entity.Person {
	po := p.personDao.FindByID(personID)
	return p.personMapper.ToDomain(po)
}

// FindByLeaderByPersonID implements facade.PersonRepositoryInterface.
func (p *PersonRepository) FindByLeaderByPersonID(personID string) *entity.Person {
	po := p.personDao.FindLeaderByPersonID(personID)
	return p.personMapper.ToDomain(po)
}
