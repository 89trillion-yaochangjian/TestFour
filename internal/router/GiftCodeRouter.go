package router

import (
	"MongoGift/internal/ctrl"
	"MongoGift/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GiftCodeRouter() {
	r := gin.Default()
	utils.InitClient()
	utils.MongoClient()
	rdb := utils.Rdb
	fmt.Println(rdb)
	r.POST("/CreateGiftCode", ctrl.CreateGiftCode)
	r.GET("/GetGiftCodeInfo", ctrl.GetGiftCodeInfoCtrl)
	r.GET("/VerifyGiftCode", ctrl.VerifyGiftCodeCtrl)
	r.GET("/UserLoginCtrl", ctrl.UserLoginCtrl)
	r.Run(":8080")
}
