package handle

import (
	"net/http"

	_ "gin/docs"
	usecaseusr "gin/internal/usecase/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files" //新版 swaggerFiles
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ApiHandler struct {
	Usecase usecaseusr.UsecaseUser
}

func NewRouter(engine *gin.Engine, usecase usecaseusr.UsecaseUser) {

	api := ApiHandler{
		Usecase: usecase,
	}

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	engine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "authorization"},
		AllowCredentials: false,
	}))

	engine.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	//engine.POST("/user", api.RegisterAccount)

	subDomain := engine.Group("/v1")
	subDomain.GET("/token", GetToken)
	subDomain.POST("/login", api.Login)
	subDomain.POST("/user", api.RegisterAccount)
	subDomain.Use(CheckAuthorization)
	subDomain.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	subDomain.POST("/test", tryTestApi)
}
