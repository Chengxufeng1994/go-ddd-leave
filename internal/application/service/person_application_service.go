package service

import (
	"github.com/Chengxufeng1994/go-ddd-leave/internal/application/usecase"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/entity"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/service"
)

type PersonApplicationService struct {
	personDomainService *service.PersonDomainService
}

func NewPersonApplicationService(personDomainService *service.PersonDomainService) usecase.PersonUseCase {
	return PersonApplicationService{}
}

// FindByID implements usecase.PersonUseCase.
func (p PersonApplicationService) FindByID(personID string) (*entity.Person, error) {
	panic("unimplemented")
}

// FindFirstApprover implements usecase.PersonUseCase.
func (p PersonApplicationService) FindFirstApprover(applicantID string, leaderMaxLevel int) (*entity.Person, error) {
	person := p.personDomainService.FindFirstApprover(applicantID, leaderMaxLevel)
	if person == nil {
		return nil, nil
	}

	return person, nil
}
