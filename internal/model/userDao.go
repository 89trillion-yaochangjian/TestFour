package model

import (
	"MongoGift/internal/status"
	"MongoGift/internal/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

//登陆

func FindUser(UId string) (User, *status.Response) {
	// 创建一个userStruct变量用来接收查询的结果
	var userStruct User
	//以为用户输入字符串为用户名
	filter := bson.D{{"uid", UId}}
	utils.MongoCon.FindOne(context.TODO(), filter).Decode(&userStruct)
	if userStruct.UID == "" {
		//Uid为空则注册用户
		userStruct.UID = utils.CreateUID()
		err1 := InsertUser(userStruct)
		if err1 != nil {
			return userStruct, err1
		}
	}
	return userStruct, nil
}

//更新用户奖励信息

func UpdateUser(user User, CodeInfo GiftCodeInfo) *status.Response {
	user.GoldCoins = CodeInfo.ContentList.GoldCoins
	user.Diamonds = CodeInfo.ContentList.Diamonds
	//以为用户输入字符串为用户名
	filter := bson.D{{"uid", user.UID}}
	//更新用户,
	update := bson.D{
		{"$inc", bson.D{
			{"goldcoins", user.GoldCoins},
			{"diamonds", user.Diamonds},
		}},
	}
	updateResult, err1 := utils.MongoCon.UpdateOne(context.TODO(), filter, update)
	if err1 != nil {
		return status.DBUpdateErr
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return nil
}

//注册用户

func InsertUser(user User) *status.Response {

	insertResult, err := utils.MongoCon.InsertOne(context.TODO(), user)
	if err != nil {
		return status.DBInsertErr
	}
	fmt.Println("Inserted a single document: ", insertResult)
	return nil
}
