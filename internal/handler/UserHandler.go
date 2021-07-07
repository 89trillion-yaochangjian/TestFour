package handler

import (
	"MongoGift/StructInfo"
	"MongoGift/internal/service"
)

//注册与登录

func UserLoginHandler(str string)  (StructInfo.MesInfo,error){
	User,err := service.UserLoginServer(str)
	if err != nil {
		return StructInfo.MesInfo{Msg: "注册与登录失败",ER: err},nil
	}
	return StructInfo.MesInfo{Msg: "注册与登录成功",Data: User},nil
}