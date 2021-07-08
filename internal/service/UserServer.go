package service

import (
	"MongoGift/internal/dao"
	"MongoGift/internal/structInfo"
	"MongoGift/internal/utils"
)

//登陆，返回用户信息

func UserLoginServer(str string) (structInfo.User, *structInfo.Response) {
	//根据UID判断是否存在该玩家
	userInfo, err := dao.FindUser(str)
	if err != nil {
		return userInfo, err
	}
	if len(userInfo.User) == 0 {
		//创建用户,调用Uid
		userInfo.UID = utils.CreateUID()
		userInfo.User = str
		err1 := dao.InsertUser(userInfo)
		if err1 != nil {
			return userInfo, err1
		}
	}
	return userInfo, nil
}
