package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary test api
// @Tags test
// @Description  try test api
// @Produce json
// @Router /test [post]
// @Param authorization header string true "bearer token"
// @Failure 500 {string} json
// @Success 200 {string} json
func tryTestApi(c *gin.Context) {
	payload, ok := c.Get("Auth_payload")

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "payload doesn't exist"})
		return
	}

	c.JSON(http.StatusOK, payload)

}
