package facade

import (
	"net/http"

	"github.com/Chengxufeng1994/go-ddd-leave/internal/application/usecase"
	infraapi "github.com/Chengxufeng1994/go-ddd-leave/internal/infrastructure/api"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/interface/assembler"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/interface/dto"
	"github.com/gin-gonic/gin"
)

type LeaveApi struct {
	leaveApplicationService usecase.LeaveUseCase
	leaveAssembler          *assembler.LeaveAssembler
}

func NewLeaveApi(leaveApplicationService usecase.LeaveUseCase) *LeaveApi {
	return &LeaveApi{
		leaveApplicationService: leaveApplicationService,
		leaveAssembler:          assembler.NewLeaveAssembler(),
	}
}

func (api *LeaveApi) CreateLeaveInfo(c *gin.Context) {
	var leaveDto dto.LeaveDTO
	if err := c.ShouldBindJSON(&leaveDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	leaveDo := api.leaveAssembler.ToDO(&leaveDto)
	api.leaveApplicationService.CreateLeaveInfo(leaveDo)
	c.JSON(http.StatusOK, infraapi.Ok(nil))
	return
}

func (api *LeaveApi) UpdateLeaveInfo(c *gin.Context) {}

func (api *LeaveApi) SubmitApproval(c *gin.Context) {}

func (api *LeaveApi) FindByID(c *gin.Context) {
	leaveID := c.Param("leaveId")
	entity, err := api.leaveApplicationService.GetLeaveInfo(leaveID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, infraapi.Ok(api.leaveAssembler.ToDTO(entity)))
}

func (api *LeaveApi) QueryByApplicantID(c *gin.Context) {
	applicatnID := c.Param("applicantId")
	entites, _ := api.leaveApplicationService.QueryLeaveInfosByApplicant(applicatnID)
	dtos := make([]*dto.LeaveDTO, 0, len(entites))
	for i := 0; i < len(entites); i++ {
		dtos = append(dtos, api.leaveAssembler.ToDTO(entites[i]))
	}
	c.JSON(http.StatusOK, infraapi.Ok(dtos))
}
