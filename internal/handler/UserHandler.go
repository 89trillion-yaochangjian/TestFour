package handler

import (
	"MongoGift/internal/service"
	"MongoGift/internal/structInfo"
)

//注册与登录

func UserLoginHandler(str string) (structInfo.User, *structInfo.Response) {
	User, err := service.UserLoginServer(str)
	if err != nil {
		return User, err
	}
	return User, nil
}
