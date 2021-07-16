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
	UId := c.Query("uid")
	if len(UId) == 0 {
		c.JSON(http.StatusBadRequest, status.StringErr)
		return
	}
	info, err := service.UserLoginServer(UId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if info.UID != UId {
		c.JSON(http.StatusOK, status.Response{Code: status.SUCCESS, Msg: status.UserADDSUCCESS, Data: info})
		return
	}
	c.JSON(http.StatusOK, status.OK.WithData(info))
}
