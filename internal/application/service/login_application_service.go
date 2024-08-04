package service

import (
	"github.com/Chengxufeng1994/go-ddd-leave/internal/application/usecase"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/entity"
)

type LoginApplicationService struct {
}

func NewLoginApplicationService() usecase.AuthUseCase {
	return &LoginApplicationService{}
}

// Login implements usecase.AuthUseCase.
func (l *LoginApplicationService) Login(*entity.Person) error {
	panic("unimplemented")
}
