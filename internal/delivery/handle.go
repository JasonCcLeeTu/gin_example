package handle

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ApiRouter(engine *gin.Engine) {

	engine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "authorization"},
		AllowCredentials: false,
	}))

	engine.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	engine.GET("/token", GetToken)
	subDomain := engine.Group("/v1").Use(CheckAuthorization)
	subDomain.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	subDomain.POST("/test", tryTestApi)

}
