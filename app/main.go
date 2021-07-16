package main

import (
	"MongoGift/internal/router"
	"MongoGift/internal/utils"
)

func main() {
	//初始化连接
	utils.InitClient()
	utils.MongoClient()
	//调用路由
	router.GiftCodeRouter()
}
