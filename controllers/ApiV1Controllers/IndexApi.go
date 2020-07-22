package ApiV1Controllers

import (
	"dgg/module"
	"dgg/module/model"
	"dgg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexRequest struct {
	Id   int64  `form:"id" valid:"required~id必须是整数"`
	Name string `form:"name" valid:"required~名字不能为空,runelength(1|10)~名字在1-10个字符之间"`
}

func Index(c *gin.Context) {

	admin := []model.AdminsModel{}

	module.DB.Offset(0).Limit(1).Find(&admin)

	c.JSON(http.StatusOK, util.GetApiJsonResult("200", "success", &admin))
}

func SetAdmin(c *gin.Context) {

	var form *IndexRequest
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}
