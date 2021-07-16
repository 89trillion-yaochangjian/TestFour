package service

import (
	"MongoGift/internal/dao"
	"MongoGift/internal/model"
	"MongoGift/internal/response"
	"MongoGift/internal/status"
	"MongoGift/internal/utils"
	"encoding/json"
	"errors"
	"github.com/golang/protobuf/proto"
	"time"
)

var receiveGiftList model.ReceiveGiftList

//管理后台调用 - 创建礼品码

func CreateGiftCodeService(giftCodeInfo model.GiftCodeInfo) (string, *status.Response) {
	code := utils.GetGiftCodeUtil()
	giftCodeInfo.Code = code
	//设置创建时间
	giftCodeInfo.CreatTime = time.Now()
	//设置过期时间
	validPeriod := giftCodeInfo.ValidPeriod
	jsonCodeInfo, err1 := json.Marshal(giftCodeInfo)
	if err1 != nil {
		return "", status.MarshalErr
	}
	CodeInfo, err := dao.CreateGiftCodeDao(code, jsonCodeInfo, validPeriod)
	if err != nil {
		return "", err
	}
	return CodeInfo, err
}

//管理后台调用 - 查询礼品码信息

func GetGiftCodeInfoService(code string) (model.GiftCodeInfo, *status.Response) {
	//根据礼品码查询礼品信息
	CodeInfo, err := dao.GetGiftCodeInfoDao(code)
	if err != nil {
		return CodeInfo, err
	}
	//显示礼包类型
	codeType := CodeInfo.CodeType
	if codeType > 0 {
		CodeInfo.CodeTypeDesc = "不指定用户限制兑换次数"
	} else if codeType == -1 {
		CodeInfo.CodeTypeDesc = "指定用户一次性消耗"
	} else if codeType == -2 {
		CodeInfo.CodeTypeDesc = "不限用户不限次数兑换"
	}
	return CodeInfo, err
}

//客户端调用 - 验证礼品码

func VerifyFiftCodeService(code string, Uid string) ([]byte, error) {
	Reward := response.GeneralReward{
		Changes: make(map[uint32]uint64),
		Balance: make(map[uint32]uint64),
		Counter: make(map[uint32]uint64),
	}
	//获取礼包码信息
	CodeInfo, GteInfoRrr := dao.GetGiftCodeInfoDao(code)
	if GteInfoRrr != nil {
		err := errors.New("礼包码无效")
		return nil, err
	}
	userInfo, _ := dao.FindUser(Uid)
	//根据礼包码，更新用户信息，返回二进制序列
	switch CodeInfo.CodeType {
	case -1:
		if CodeInfo.User != Uid {
			err := errors.New("指定用户领取")
			return nil, err
		}
		if CodeInfo.ReceiveNum == 1 {
			err := errors.New("您已领取，不要重复领取")
			return nil, err
		}
		Reward1, _ := dao.VerifyFiftCodeDao(CodeInfo, userInfo, Uid)
		return proto.Marshal(&Reward1)
	case 0:
		if CodeInfo.AvailableTimes > CodeInfo.ReceiveNum {
			Reward2, _ := dao.VerifyFiftCodeDao(CodeInfo, userInfo, Uid)
			return proto.Marshal(&Reward2)
		}
	case -2:
		Reward2, _ := dao.VerifyFiftCodeDao(CodeInfo, userInfo, Uid)
		return proto.Marshal(&Reward2)
	}
	return proto.Marshal(&Reward)
}
