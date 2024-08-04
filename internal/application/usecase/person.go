package usecase

import "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/entity"

type PersonUseCase interface {
	FindByID(personID string) (*entity.Person, error)
	FindFirstApprover(ApplicantID string, leaderMaxLevel int) (*entity.Person, error)
}
