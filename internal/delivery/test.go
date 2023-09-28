package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func tryTestApi(c *gin.Context) {
	payload, ok := c.Get("Auth_payload")

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "payload doesn't exist"})
		return
	}

	c.JSON(http.StatusOK, payload)

}
