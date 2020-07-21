package ApiV1Controllers

import (
	"github.com/gin-gonic/gin"
	"lidongheDemo/models"
	"lidongheDemo/models/model"
	"lidongheDemo/util"
	"net/http"
)

type IndexRequest struct {
	Id   int64  `json:"id" valid:"required~id必须是整数"`
	Name string `json:"name" valid:"required~名字不能为空,runelength(1|10)~名字在1-10个字符之间"`
}

func Index(c *gin.Context) {
	admins := &model.AdminsModel{}
	models.DB.First(&admins)

	var Name = models.GetStringValue("name")

	c.JSON(http.StatusOK, util.GetApiJsonResult("200", "success", Name))
}
