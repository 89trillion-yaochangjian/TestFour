package ctrl

import (
	"MongoGift/StructInfo"
	"MongoGift/internal/handler"
	"MongoGift/internal/response"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateGiftCode(c *gin.Context) {
	//获取参数
	info,err1 := c.GetRawData()
	if err1 != nil{
		c.JSON(http.StatusOK,StructInfo.MesInfo{Msg: "获取参数失败",Data: err1})
	}
	var giftCodeInfo StructInfo.GiftCodeInfo
	json.Unmarshal(info,&giftCodeInfo)
	//var giftCodeInfo = StructInfo.GiftCodeInfo{}
	//c.ShouldBind(&giftCodeInfo)
	//调用Handler
	code,err := handler.CreateGiftCodeHandler(giftCodeInfo)
	if err != nil {
		c.JSON(http.StatusOK,StructInfo.MesInfo{Msg: "创建礼包码失败",ER: err})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
	})
}

//管理后台调用 - 查询礼品码信息

func GetGiftCodeInfoCtrl(c *gin.Context) {
	//获取参数
	code := c.Query("code")
	info,err := handler.GetFiftCodeInfoHandler(code)
	if err != nil {
		c.JSON(http.StatusOK,StructInfo.MesInfo{Msg: "查询礼品码失败",ER: err})
	}
	c.JSON(http.StatusOK, info)
}

func VerifyGiftCodeCtrl(c *gin.Context) {
	//获取参数
	code := c.Query("code")
	user := c.Query("user")
	info,e := handler.VerifyFiftCodeHandler(code,user)
	if e!=nil {
		fmt.Printf("")
	}
	Reward := response.GeneralReward{}
	json.Unmarshal(info,&Reward)
	c.JSON(http.StatusOK,info)
}