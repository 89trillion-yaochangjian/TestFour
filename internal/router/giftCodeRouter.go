package router

import (
	"MongoGift/internal/ctrl"
	"github.com/gin-gonic/gin"
)

func GiftCodeRouter() {
	r := gin.Default()
	r.POST("/CreateGiftCode", ctrl.CreateGiftCode)
	r.GET("/GetGiftCodeInfo", ctrl.GetGiftCodeInfoCtrl)
	r.GET("/VerifyGiftCode", ctrl.VerifyGiftCodeCtrl)
	r.GET("/UserLoginCtrl", ctrl.UserLoginCtrl)
	r.Run(":8080")
}
