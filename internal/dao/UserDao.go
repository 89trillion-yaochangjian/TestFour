package dao

import (
	"MongoGift/StructInfo"
	"MongoGift/internal/utils"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

//登陆
func FindUser(userName string) (StructInfo.User,error){
	// 创建一个userStruct变量用来接收查询的结果
	var userStruct StructInfo.User
	//以为用户输入字符串为用户名
	filter := bson.D{{"user",userName}}
	err := utils.MongoCon.FindOne(context.TODO(), filter).Decode(&userStruct)
	if err!=nil {
		err = errors.New("mongodb查询用户失败")
		return userStruct,nil
	}
	fmt.Printf("Found a single document: %+v\n", userStruct)
	return userStruct,nil
}

//更新用户奖励信息
func UpdateUser(user StructInfo.User) {
	//以为用户输入字符串为用户名
	filter := bson.D{{"user",user.User}}
	//更新用户,
	update := bson.D{
		{"$inc", bson.D{
			{"goldcoins", user.GoldCoins},
			{"diamonds", user.Diamonds},
		}},
	}
	updateResult, err1 := utils.MongoCon.UpdateOne(context.TODO(), filter, update)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}



//注册用户
func InsertUser(user StructInfo.User)  {

	insertResult, err := utils.MongoCon.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult)
}