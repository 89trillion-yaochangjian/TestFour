package model

import (
	"MongoGift/internal/config"
	"MongoGift/internal/response"
	"MongoGift/internal/status"
	"encoding/json"
	"time"
)

var receiveGiftList ReceiveGiftList

//管理后台调用 - 创建礼品码

func CreateGiftCodeDao(code string, jsonCodeInfo []byte, validPeriod int) (string, *status.Response) {
	//以礼品吗为key存到Redis,并设置过期时间
	err := config.Rdb.Set(code, jsonCodeInfo, time.Duration(validPeriod)*time.Hour).Err()
	if err != nil {
		return "", status.RedisErr
	}
	return code, nil
}

//管理后台调用 - 查询礼品码信息

func GetGiftCodeInfoDao(code string) (GiftCodeInfo, *status.Response) {

	CodeInfo := GiftCodeInfo{}
	//根据礼品码查询礼品信息
	JsonCodeInfo, err1 := config.Rdb.Get(code).Result()
	if err1 != nil {
		return CodeInfo, status.RedisErr
	}
	//反序列化
	err := json.Unmarshal([]byte(JsonCodeInfo), &CodeInfo)
	if err != nil {
		return CodeInfo, status.MarshalErr
	}
	return CodeInfo, nil
}

//客户端调用 - 验证礼品码

func VerifyFiftCodeDao(giftCodeInfo GiftCodeInfo, userInfo User, Uid string) (response.GeneralReward, *status.Response) {
	// 开启一个TxPipeline事务
	pipe := config.Rdb.TxPipeline()
	defer pipe.Close()
	jsonCodeInfo, err1 := json.Marshal(giftCodeInfo)
	//领取数加一,单独存储key
	count := config.Rdb.Incr(giftCodeInfo.Code + "counts")
	giftCodeInfo.ReceiveNum = count.Val()
	//用户添加到领取列表，保存到Redis
	receiveGiftList.ReceiveTime = time.Now().Unix()
	receiveGiftList.ReceiveUser = Uid
	giftCodeInfo.ReceiveList = append(giftCodeInfo.ReceiveList, receiveGiftList)
	code := giftCodeInfo.Code
	if err1 != nil {
		return response.GeneralReward{}, status.MarshalErr
	}
	// 通过Exec函数提交redis事务
	pipe.Set(code, jsonCodeInfo, config.Rdb.TTL(code).Val())
	_, err := pipe.Exec()
	if err != nil {
		return response.GeneralReward{}, status.RedisErr
	}
	Reward, err0 := UpdateUser(userInfo, giftCodeInfo)
	if err0 != nil {
		return Reward, err0
	}
	return Reward, nil
}
