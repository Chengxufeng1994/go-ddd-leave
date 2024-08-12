package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/Chengxufeng1994/go-ddd-leave/internal/application/event"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/application/service"
	commonevt "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/common/event"
	leaveevt "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/event"
	leavedao "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/repository/dao"
	leavepersistence "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/repository/persistence"
	leavepo "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/repository/po"
	leavedomainservice "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/leave/service"
	persondao "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/repository/dao"
	personpersistence "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/repository/persistence"
	personpo "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/repository/po"
	persondomainservice "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/service"
	ruledao "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/rule/repository/dao"
	ruelpersistence "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/rule/repository/persistence"
	rulepo "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/rule/repository/po"
	ruledomainservice "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/rule/service"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/infrastructure/broker"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/infrastructure/eventstoredb"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/infrastructure/util"
	"github.com/Chengxufeng1994/go-ddd-leave/internal/interface/facade"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	gormDb, _ := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=P@ssw0rd host=10.1.5.7 dbname=postgres port=31970 sslmode=disable search_path=leave_sample TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: newLogger,
	})
	gormDb.AutoMigrate(&personpo.Person{}, &personpo.Relationship{})
	gormDb.AutoMigrate(&leavepo.LeaveEvent{}, &leavepo.Leave{}, &leavepo.ApprovalInfo{})
	gormDb.AutoMigrate(&rulepo.ApprovalRule{})

	esDb, _ := eventstoredb.NewEventStoreDB()
	defer esDb.Close()

	sendEmailHandler := event.NewSendEmailHandler()
	leavePublisher := broker.NewLeavePublisher()
	leavePublisher.Subscribe(sendEmailHandler,
		[]commonevt.Event{&leaveevt.LeaveCreatedEvent{}}...)

	sf, _ := util.NewSnowFlake()
	approvalInfoDao := leavedao.NewApprovalInfoDao(gormDb)
	leaveDao := leavedao.NewLeaveDao(gormDb)
	leaveEventDao := leavedao.NewLeaveEventDao(gormDb)
	leaveRepository := leavepersistence.NewLeaveRepository(leaveDao, approvalInfoDao, leaveEventDao, sf)
	leaveFactory := leavedomainservice.NewLeaveFactory()

	personDao := persondao.NewPersonDao(gormDb)
	personRepository := personpersistence.NewPersonRepository(personDao)
	personDomainService := persondomainservice.NewPersonDomainService(personRepository)

	approvalRuleDao := ruledao.NewApprovalRuleDao(gormDb)
	approvalRuleRepository := ruelpersistence.NewApprovalRuleRepository(approvalRuleDao)
	approvalRuleDomainService := ruledomainservice.NewApprovalRuleDomainService(approvalRuleRepository)

	leaveDomainService := leavedomainservice.NewLeaveDomainService(leaveRepository, leaveFactory, leavePublisher)
	leaveApplicationService := service.NewLeaveApplicationService(leaveDomainService, personDomainService, approvalRuleDomainService)
	loginApplicationService := service.NewLoginApplicationService()

	leaveApi := facade.NewLeaveApi(leaveApplicationService)
	authApi := facade.NewAuthApi(loginApplicationService)

	r := gin.Default()
	authGroup := r.Group("/auth")
	authGroup.POST("/login", authApi.Login)

	leaveGroup := r.Group("/leave")
	leaveGroup.POST("/", leaveApi.CreateLeaveInfo)
	leaveGroup.PUT("/", leaveApi.UpdateLeaveInfo)
	leaveGroup.POST("/submit", leaveApi.SubmitApproval)
	leaveGroup.GET("/:leaveId", leaveApi.FindByID)
	leaveGroup.GET("/query/applicant/:applicantId", leaveApi.QueryByApplicantID)
	// leaveGroup.GET("/query/approver/:approverId", leaveApi.QueryByApproverID)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
