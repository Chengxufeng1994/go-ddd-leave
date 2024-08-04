package facade

import (
	"net/http"

	"github.com/Chengxufeng1994/go-ddd-leave/internal/application/usecase"
	infraapi "github.com/Chengxufeng1994/go-ddd-leave/internal/infrastructure/api"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/interface/assembler"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/interface/dto"
	"github.com/gin-gonic/gin"
)

type AuthApi struct {
	loginApplication usecase.AuthUseCase
	personAssembler  *assembler.PersonAssembler
}

func NewAuthApi(loginApplication usecase.AuthUseCase) *AuthApi {
	return &AuthApi{
		loginApplication: loginApplication,
		personAssembler:  assembler.NewPersonAssembler(),
	}
}

func (api *AuthApi) Login(c *gin.Context) {
	var person dto.PersonDTO
	if err := c.ShouldBindJSON(&person); err != nil {
		return
	}

	if err := api.loginApplication.Login(api.personAssembler.ToDO(&person)); err != nil {
		return
	}

	c.JSON(http.StatusOK, infraapi.Ok(nil))
}
