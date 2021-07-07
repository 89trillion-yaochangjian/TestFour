package service

import (
	"MongoGift/StructInfo"
	"MongoGift/internal/dao"
	"MongoGift/internal/utils"
	"errors"
)

//注册与登录

func UserLoginServer(str string) (StructInfo.User,error){
	//根据UID判断是否存在该玩家
	userInfo,err := dao.FindUser(str)
	if err!=nil {
		err = errors.New("查询用户失败")
	}
	if len(userInfo.UID)==0{
		//创建用户,调用Uid
		userInfo.UID = utils.CreateUID()
		userInfo.User = str
		dao.InsertUser(userInfo)
	}
	return userInfo,nil
}

