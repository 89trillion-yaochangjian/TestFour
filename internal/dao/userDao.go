package dao

import (
	"MongoGift/internal/model"
	"MongoGift/internal/status"
	"MongoGift/internal/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

//登陆

func FindUser(userName string) (model.User, *status.Response) {
	// 创建一个userStruct变量用来接收查询的结果
	var userStruct model.User
	//以为用户输入字符串为用户名
	filter := bson.D{{"user", userName}}
	err := utils.MongoCon.FindOne(context.TODO(), filter).Decode(&userStruct)
	if err != nil {
		return userStruct, status.LoginUserErr
	}
	fmt.Printf("Found a single document: %+v\n", userStruct)
	return userStruct, nil
}

//更新用户奖励信息

func UpdateUser(user model.User, CodeInfo model.GiftCodeInfo) *status.Response {
	user.GoldCoins = CodeInfo.ContentList.GoldCoins
	user.Diamonds = CodeInfo.ContentList.Diamonds
	//以为用户输入字符串为用户名
	filter := bson.D{{"user", user.User}}
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

func InsertUser(user model.User) *status.Response {

	insertResult, err := utils.MongoCon.InsertOne(context.TODO(), user)
	if err != nil {
		return status.DBInsertErr
	}
	fmt.Println("Inserted a single document: ", insertResult)
	return nil
}