package app

import (
	_ "database/sql"
	"log/slog"

	auth2 "hackathons-app/internal/api/http/v1/auth"
	httpAuth "hackathons-app/internal/api/http/v1/auth"
	httpHackathon "hackathons-app/internal/api/http/v1/hackathon"
	httpUser "hackathons-app/internal/api/http/v1/user"
	"hackathons-app/internal/config"
	"hackathons-app/internal/db"
	"hackathons-app/internal/models"
	"hackathons-app/internal/repositories"
	"hackathons-app/internal/services"

	_ "hackathons-app/docs"

	"hackathons-app/pkg/log"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Run @title Go API
//
//	@version					1.0
//	@description				This is a sample server.
//	@termsOfService				http://swagger.io/terms/
//	@contact.name				API Support
//	@contact.url				http://www.swagger.io/support
//	@contact.email				support@swagger.io
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//	@host						localhost:4242
//	@securitydefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Enter the token with the `Bearer ` prefix, e.g. "Bearer abcde12345".
func Run() {
	appConfig, gormConfig := config.GetConfig(".env"), gorm.Config{Logger: logger.Default.LogMode(logger.Info)}
	baseLogger := log.NewLogger().With("app", "main")
	bdLogger := baseLogger.With("component", "database")

	bdLogger.Info("", slog.Any("dsn = ", appConfig.DB.GetDsn()))
	dialector := postgres.Open(
		appConfig.
			DB.GetDsn(),
	)

	gormDb, err := db.NewGormDatabase(dialector, &gormConfig)
	if err != nil {
		bdLogger.Error(
			"Error connecting to database: ",
			slog.Any("err", err),
			slog.Any("config", appConfig.DB),
		)

		panic(err)
	}
	bdLogger.Info("Success Connected to database")

	bdLogger.Info("Starting Migrations")
	err = gormDb.GetDB().AutoMigrate(&models.User{}, &models.Hackathon{})
	if err != nil {
		bdLogger.Error("Error migrating database: ", slog.Any("err", err))
		panic(err)
	}
	bdLogger.Info("Success Migrated database")

	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	var (
		userRepo         = repositories.NewUserRepository(gormDb)
		hackathonRepo    = repositories.NewHackathonRepository(gormDb)
		userService      = services.NewUserService(userRepo, hackathonRepo)
		hackathonService = services.NewHackathonService(hackathonRepo)
		hackathonHandler = httpHackathon.NewHackathonHandler(hackathonService)
		auth             = auth2.NewAuth(baseLogger, appConfig.Jwt, userService)
		authHandler      = httpAuth.NewAuthHandler(userService, appConfig.Jwt)
		userHandler      = httpUser.NewUserHandler(userService)
	)
	httpHackathon.RegisterRouters(r, hackathonHandler, auth)
	httpUser.RegisterRouters(r, userHandler, auth)
	httpAuth.RegisterRoutes(r, authHandler)

	err = r.Run(appConfig.Server.GetURL())
	if err != nil {
		panic(err)
	}
}
