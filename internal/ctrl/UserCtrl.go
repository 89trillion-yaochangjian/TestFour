package ctrl

import (
	"MongoGift/internal/handler"
	"MongoGift/internal/structInfo"
	"github.com/gin-gonic/gin"
	"net/http"
)

//玩家注册登陆

func UserLoginCtrl(c *gin.Context) {
	//获取参数
	user := c.Query("str")
	if len(user) == 0 {
		c.JSON(http.StatusBadRequest, structInfo.StringErr)
	}
	info, err := handler.UserLoginHandler(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, structInfo.OK.WithData(info))
}
