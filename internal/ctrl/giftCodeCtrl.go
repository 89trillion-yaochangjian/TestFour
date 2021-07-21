package ctrl

import (
	"MongoGift/internal/model"
	"MongoGift/internal/response"
	"MongoGift/internal/service"
	"MongoGift/internal/status"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

//管理后台调用 - 创建礼品码

func CreateGiftCode(c *gin.Context) {
	//获取参数
	info, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, status.ParamErr)
		return
	}
	var giftCodeInfo model.GiftCodeInfo
	json.Unmarshal(info, &giftCodeInfo)
	// 0--不限定用户，限定领取次数   -1--指定用户一次领取  -2--不限定用户，不限定次数
	if giftCodeInfo.CodeType != -1 && giftCodeInfo.CodeType != 0 && giftCodeInfo.CodeType != -2 {
		c.JSON(http.StatusBadRequest, status.CodeTypeErr)
		return
	}
	//指定用户一次领取参数判断
	if giftCodeInfo.CodeType == -1 && len(giftCodeInfo.User) == 0 {
		c.JSON(http.StatusBadRequest, status.CodeUIDErr)
		return
	}

	code, err1 := service.CreateGiftCodeService(giftCodeInfo)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, err1)
		return
	}
	c.JSON(http.StatusOK, status.OK.WithData(code))
}

//管理后台调用 - 查询礼品码信息

func GetGiftCodeInfoCtrl(c *gin.Context) {
	//获取参数
	code := c.Query("code")
	if len(code) != 8 {
		c.JSON(http.StatusBadRequest, status.CodeLenErr)
		return
	}
	info, err := service.GetGiftCodeInfoService(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, status.OK.WithData(info))
}

//客户端调用 - 验证礼品码

func VerifyGiftCodeCtrl(c *gin.Context) {
	//获取参数
	code := c.Query("code")
	if len(code) != 8 {
		c.JSON(http.StatusBadRequest, status.CodeLenErr)
		return
	}
	Uid := c.Query("uid")
	if len(Uid) == 0 {
		c.JSON(http.StatusBadRequest, status.CodeUIDErr)
		return
	}
	info, err := service.VerifyFiftCodeService(code, Uid)
	if err != nil {
		switch err.Error() {
		case "礼包码无效":
			c.JSON(http.StatusBadRequest, status.CodeErr)
			return
		case "指定用户领取":
			c.JSON(http.StatusBadRequest, status.OrderUser)
			return
		case "您已领取，不要重复领取":
			c.JSON(http.StatusBadRequest, status.GetCodeSecond)
			return
		}
	}
	Reward := response.GeneralReward{}
	json.Unmarshal(info, &Reward)
	c.JSON(http.StatusOK, status.OK.WithData(info))
}
