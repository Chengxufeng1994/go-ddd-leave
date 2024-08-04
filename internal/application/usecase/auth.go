package usecase

import "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/entity"

type AuthUseCase interface {
	Login(*entity.Person) error
}
