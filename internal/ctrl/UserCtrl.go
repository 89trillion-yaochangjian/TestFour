package ctrl

import (
	"MongoGift/StructInfo"
	"MongoGift/internal/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLoginCtrl(c *gin.Context) {
	//获取参数
	user := c.Query("str")
	info,err := handler.UserLoginHandler(user)
	if err != nil {
		c.JSON(http.StatusOK,StructInfo.MesInfo{Msg: "注册登陆失败",ER: err})
	}
	c.JSON(http.StatusOK, info)
}

