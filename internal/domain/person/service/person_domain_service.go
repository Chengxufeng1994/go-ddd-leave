package service

import (
	"time"

	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/entity"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/repository/facade"
)

type PersonDomainService struct {
	personRepository facade.PersonRepositoryInterface
}

func NewPersonDomainService(personRepository facade.PersonRepositoryInterface) *PersonDomainService {
	return &PersonDomainService{
		personRepository: personRepository,
	}
}

func (personDomainService *PersonDomainService) Create(person *entity.Person) {
	existed := personDomainService.personRepository.FindByID(person.PersonID)
	if existed != nil {
		return
	}

	person.Create()
	personDomainService.personRepository.Save(person)
}

func (personDomainService *PersonDomainService) Update(person *entity.Person) {
	person.LastModifiedTime = time.Now()
	personDomainService.personRepository.Update(person)
}

func (personDomainService *PersonDomainService) DeleteByID(personID string) {
	person := personDomainService.personRepository.FindByID(personID)
	person.Disable()
	personDomainService.personRepository.Update(person)
}

func (personDomainService *PersonDomainService) FindFirstApprover(applicantID string, leaderMaxLevel int) *entity.Person {
	leader := personDomainService.personRepository.FindByLeaderByPersonID(applicantID)
	if leader.RoleLevel > leaderMaxLevel {
		return nil
	}

	return leader
}

func (personDomainService *PersonDomainService) FindNextApprover(currentApproverID string, leaderMaxLevel int) *entity.Person {
	leader := personDomainService.personRepository.FindByLeaderByPersonID(currentApproverID)
	if leader.RoleLevel > leaderMaxLevel {
		return nil
	}

	return leader
}
