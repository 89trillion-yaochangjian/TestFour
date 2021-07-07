package service

import (
	"MongoGift/StructInfo"
	"MongoGift/internal/dao"
	"MongoGift/internal/response"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)
var receiveGiftList StructInfo.ReceiveGiftList
//管理后台调用 - 创建礼品码

func CreateGiftCodeService(giftCodeInfo StructInfo.GiftCodeInfo) (string,error) {
	CodeInfo,err := dao.CreateGiftCodeDao(giftCodeInfo)
	if err!=nil {
		err = errors.New("创建礼品码异常")
		return "",err
	}
	return CodeInfo,err
}


//管理后台调用 - 查询礼品码信息

func GetGiftCodeInfoService(code string) (StructInfo.GiftCodeInfo,error){
	//根据礼品码查询礼品信息
	CodeInfo,err := dao.GetGiftCodeInfoDao(code)
	if err!=nil {
		err = errors.New("创建礼品码异常")
		return CodeInfo,err
	}
	//显示礼包类型
	codeType :=CodeInfo.CodeType
	if codeType > 0 {
		CodeInfo.CodeTypeDesc = "不指定用户限制兑换次数"
	}else if codeType == -1 {
		CodeInfo.CodeTypeDesc = "指定用户一次性消耗"
	}else if codeType == -2 {
		CodeInfo.CodeTypeDesc = "不限用户不限次数兑换"
	}
    return CodeInfo,err
}


//客户端调用 - 验证礼品码

func VerifyFiftCodeService(code string, userName string) ([]byte, error) {
	Reward := response.GeneralReward{
		Changes:make(map[uint32]uint64),
		Balance:make(map[uint32]uint64),
		Counter:make(map[uint32]uint64),
	}
	//获取礼包码信息
	CodeInfo, GteInfoRrr :=dao.GetGiftCodeInfoDao(code)
	if GteInfoRrr != nil {
		fmt.Printf("code无效或已过期, err:%v\n", GteInfoRrr)
	}
	userInfo,_ := dao.FindUser(userName)
	//获取当前客户
	switch CodeInfo.CodeType {
	case -1:
		if CodeInfo.ReceiveNum == 1||CodeInfo.User!=userName{
			fmt.Printf("礼包已经领取过")
		}
	case 0:
		if CodeInfo.AvailableTimes>CodeInfo.ReceiveNum{
			//领取数加一
			CodeInfo.ReceiveNum = CodeInfo.ReceiveNum + 1
			//用户添加到领取列表，保存到Redis
			receiveGiftList.ReceiveTime = time.Now()
			receiveGiftList.ReceiveUser = userName
			//更新用户奖励数量，保存到Mongodb
			Reward.Changes[1]=uint64(CodeInfo.ContentList.GoldCoins)
			Reward.Changes[2]=uint64(CodeInfo.ContentList.Diamonds)
			userInfo.GoldCoins = CodeInfo.ContentList.GoldCoins
			userInfo.Diamonds = CodeInfo.ContentList.Diamonds
			dao.UpdateUser(userInfo)
			Reward.Balance[1]=uint64(userInfo.GoldCoins+CodeInfo.ContentList.GoldCoins)
			Reward.Balance[2]=uint64(userInfo.Diamonds+CodeInfo.ContentList.Diamonds)
			Reward.Counter[1]=uint64(userInfo.GoldCoins+CodeInfo.ContentList.GoldCoins)
			Reward.Counter[2]=uint64(userInfo.Diamonds+CodeInfo.ContentList.Diamonds)
			CodeInfo.ReceiveList = append(CodeInfo.ReceiveList,receiveGiftList)
			dao.VerifyFiftCodeDao(CodeInfo)

		}
	case -2:
		//领取数加一
		CodeInfo.ReceiveNum = CodeInfo.ReceiveNum + 1
		//用户添加到领取列表，保存到Redis
		receiveGiftList.ReceiveTime = time.Now()
		receiveGiftList.ReceiveUser = userName
		//更新用户奖励数量，保存到Mongodb
		//金币ID为1，钻石ID为2
		Reward.Changes[1]=uint64(CodeInfo.ContentList.GoldCoins)
		Reward.Changes[2]=uint64(CodeInfo.ContentList.Diamonds)
		userInfo.GoldCoins = CodeInfo.ContentList.GoldCoins
		userInfo.Diamonds = CodeInfo.ContentList.Diamonds
		dao.UpdateUser(userInfo)
		Reward.Balance[1]=uint64(userInfo.GoldCoins+CodeInfo.ContentList.GoldCoins)
		Reward.Balance[2]=uint64(userInfo.Diamonds+CodeInfo.ContentList.Diamonds)
		Reward.Counter[1]=uint64(userInfo.GoldCoins)
		Reward.Counter[2]=uint64(userInfo.Diamonds)
		CodeInfo.ReceiveList = append(CodeInfo.ReceiveList,receiveGiftList)
		dao.VerifyFiftCodeDao(CodeInfo)
	}
	return json.Marshal(&Reward)
}