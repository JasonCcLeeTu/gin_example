package handle

import (
	"net/http"
	"time"

	entity_jwt "gin/internal/entities/jwt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetToken(c *gin.Context) {

	newtoken := entity_jwt.NewJwtAnalysis()

	usrId, err := uuid.NewRandom()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	t, err := newtoken.CreateToken(usrId, "Jason", time.Now().Add(time.Hour))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"token": t})

}

func CheckAuthorization(c *gin.Context) {
	auth := c.GetHeader("authorization")

	token, err := entity_jwt.CheckTokenFormat(auth)
	if err != nil {

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jwtAnalysis := entity_jwt.NewJwtAnalysis()

	payload, err := jwtAnalysis.VerifyToken(token)

	if err != nil {

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Set("Auth_payload", payload)
	c.Next()

}
