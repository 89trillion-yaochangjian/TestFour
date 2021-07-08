package dao

import (
	"MongoGift/internal/response"
	"MongoGift/internal/structInfo"
	"MongoGift/internal/utils"
	"encoding/json"
	"time"
)

var receiveGiftList structInfo.ReceiveGiftList

//管理后台调用 - 创建礼品码

func CreateGiftCodeDao(code string, jsonCodeInfo []byte, validPeriod int) (string, *structInfo.Response) {
	//以礼品吗为key存到Redis,并设置过期时间
	err := utils.Rdb.Set(code, jsonCodeInfo, time.Duration(validPeriod)*time.Hour).Err()
	if err != nil {
		return "", structInfo.RedisErr
	}
	return code, nil
}

//管理后台调用 - 查询礼品码信息

func GetGiftCodeInfoDao(code string) (structInfo.GiftCodeInfo, *structInfo.Response) {

	CodeInfo := structInfo.GiftCodeInfo{}
	//根据礼品码查询礼品信息
	JsonCodeInfo, err1 := utils.Rdb.Get(code).Result()
	if err1 != nil {
		return CodeInfo, structInfo.RedisErr
	}
	//反序列化
	err := json.Unmarshal([]byte(JsonCodeInfo), &CodeInfo)
	if err != nil {
		return CodeInfo, structInfo.MarshalErr
	}
	return CodeInfo, nil
}

//客户端调用 - 验证礼品码

func VerifyFiftCodeDao(giftCodeInfo structInfo.GiftCodeInfo, userInfo structInfo.User, user string) (response.GeneralReward, *structInfo.Response) {
	Reward := response.GeneralReward{
		Changes: make(map[uint32]uint64),
		Balance: make(map[uint32]uint64),
		Counter: make(map[uint32]uint64),
	}
	//更新用户奖励数量，保存到Mongodb
	//金币ID为1，钻石ID为2
	Reward.Changes[1] = uint64(giftCodeInfo.ContentList.GoldCoins)
	Reward.Changes[2] = uint64(giftCodeInfo.ContentList.Diamonds)
	userInfo.GoldCoins = giftCodeInfo.ContentList.GoldCoins
	userInfo.Diamonds = giftCodeInfo.ContentList.Diamonds
	UpdateUser(userInfo, giftCodeInfo)
	Reward.Balance[1] = uint64(userInfo.GoldCoins + giftCodeInfo.ContentList.GoldCoins)
	Reward.Balance[2] = uint64(userInfo.Diamonds + giftCodeInfo.ContentList.Diamonds)
	Reward.Counter[1] = uint64(userInfo.GoldCoins + giftCodeInfo.ContentList.GoldCoins)
	Reward.Counter[2] = uint64(userInfo.Diamonds + giftCodeInfo.ContentList.Diamonds)
	giftCodeInfo.ReceiveList = append(giftCodeInfo.ReceiveList, receiveGiftList)

	//领取数加一
	giftCodeInfo.ReceiveNum = giftCodeInfo.ReceiveNum + 1
	//用户添加到领取列表，保存到Redis
	receiveGiftList.ReceiveTime = time.Now()
	receiveGiftList.ReceiveUser = user
	giftCodeInfo.ReceiveList = append(giftCodeInfo.ReceiveList, receiveGiftList)
	code := giftCodeInfo.Code
	jsonCodeInfo, err1 := json.Marshal(giftCodeInfo)
	if err1 != nil {
		return Reward, structInfo.MarshalErr
	}
	err := utils.Rdb.Set(code, jsonCodeInfo, utils.Rdb.TTL(code).Val())
	if err != nil {
		return Reward, structInfo.RedisErr
	}
	return Reward, nil
}
