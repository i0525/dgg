package v1

import (
	"dgg/app/validate"
	"dgg/module"
	"dgg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetAdmin(c *gin.Context) {

	var login validate.Login

	err := c.ShouldBind(&login)

	if err != nil {
		c.JSON(http.StatusOK, util.GetApiJsonErrResult( module.Translate(err)))
		return
	}

	c.JSON(http.StatusOK, util.GetApiJsonResult("200", "success", "处理成功"))
}