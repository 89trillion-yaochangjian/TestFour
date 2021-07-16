package service

import (
	"MongoGift/internal/dao"
	"MongoGift/internal/model"
	"MongoGift/internal/status"
	"MongoGift/internal/utils"
)

//登陆，返回用户信息

func UserLoginServer(UId string) (model.User, *status.Response) {
	//根据UID判断是否存在该玩家
	userInfo, err := dao.FindUser(UId)
	if err != nil {
		return userInfo, err
	}
	if len(userInfo.UID) == 0 {
		//Uid为空注册用户
		userInfo.UID = utils.CreateUID()
		err1 := dao.InsertUser(userInfo)
		if err1 != nil {
			return userInfo, err1
		}
		return userInfo, status.UserADD
	}
	return userInfo, nil
}
