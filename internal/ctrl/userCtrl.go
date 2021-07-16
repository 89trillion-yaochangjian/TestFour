package ctrl

import (
	"MongoGift/internal/service"
	"MongoGift/internal/status"
	"github.com/gin-gonic/gin"
	"net/http"
)

//玩家注册登陆

func UserLoginCtrl(c *gin.Context) {
	//获取参数
	user := c.Query("str")
	if len(user) == 0 {
		c.JSON(http.StatusBadRequest, status.StringErr)
	}
	info, err := service.UserLoginServer(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, status.OK.WithData(info))
}
