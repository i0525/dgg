package v1

import (
	"dgg/app/validate"

	"github.com/gin-gonic/gin"
	"net/http"
)

func SetAdmin(c *gin.Context) {
	var form *validate.IndexRequest
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}