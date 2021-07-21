package model

import (
	"MongoGift/internal/config"
	"MongoGift/internal/response"
	"MongoGift/internal/status"
	"MongoGift/internal/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//登陆

func FindUser(UId string) (User, *status.Response) {
	// 创建一个userStruct变量用来接收查询的结果
	var userStruct User
	//以为用户输入字符串为用户名
	filter := bson.D{{"uid", UId}}
	config.MongoCon.FindOne(context.TODO(), filter).Decode(&userStruct)
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

func UpdateUser(user User, CodeInfo GiftCodeInfo) (response.GeneralReward, *status.Response) {
	Reward := response.GeneralReward{
		Changes: make(map[uint32]uint64),
		Balance: make(map[uint32]uint64),
		Counter: make(map[uint32]uint64),
	}
	if config.MongoCon == nil {
		return Reward, status.MongoDBErr
	}
	//更新用户奖励数量，保存到Mongodb
	//金币ID为1，钻石ID为2
	Reward.Changes[1] = uint64(CodeInfo.ContentList.GoldCoins)
	Reward.Changes[2] = uint64(CodeInfo.ContentList.Diamonds)
	Reward.Balance[1] = uint64(user.GoldCoins + CodeInfo.ContentList.GoldCoins)
	Reward.Balance[2] = uint64(user.Diamonds + CodeInfo.ContentList.Diamonds)
	Reward.Counter[1] = uint64(user.GoldCoins + CodeInfo.ContentList.GoldCoins)
	Reward.Counter[2] = uint64(user.Diamonds + CodeInfo.ContentList.Diamonds)

	//开启是事务
	err := config.Session.StartTransaction()
	if err != nil {
		return Reward, status.MongoDBTractionErr
	}
	monCtx := mongo.NewSessionContext(context.TODO(), config.Session)
	//以为用户输入字符串为用户唯一识别
	filter := bson.D{{"uid", user.UID}}
	//更新用户,
	update := bson.D{
		{"$inc", bson.D{
			{"goldcoins", CodeInfo.ContentList.GoldCoins},
			{"diamonds", CodeInfo.ContentList.Diamonds},
		}},
	}
	_, err1 := config.MongoCon.UpdateOne(monCtx, filter, update)
	if err1 != nil {
		config.Session.AbortTransaction(context.TODO())
		return Reward, status.DBUpdateErr
	}
	config.Session.CommitTransaction(context.TODO())
	return Reward, nil
}

//注册用户

func InsertUser(user User) *status.Response {

	insertResult, err := config.MongoCon.InsertOne(context.TODO(), user)
	if err != nil {
		return status.DBInsertErr
	}
	fmt.Println("Inserted a single document: ", insertResult)
	return nil
}
