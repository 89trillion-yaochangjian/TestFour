package main

import (
	"MongoGift/internal/config"
	"MongoGift/internal/router"
)

func main() {
	//初始化连接
	config.InitClient()
	config.MongoClient()
	//调用路由
	router.GiftCodeRouter()
}
