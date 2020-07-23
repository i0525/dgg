package v1

import (
	"dgg/app/model"
	"dgg/module"
	"dgg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {

	admin := []model.AdminsModel{}

	module.DB.Offset(0).Limit(1).Find(&admin)

	c.JSON(http.StatusOK, util.GetApiJsonResult("200", "success", &admin))
}


