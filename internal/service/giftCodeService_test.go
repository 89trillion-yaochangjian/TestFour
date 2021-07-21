package service

import (
	"MongoGift/internal/config"
	structInfo2 "MongoGift/internal/model"
	"MongoGift/internal/response"
	"fmt"
	"github.com/golang/protobuf/proto"
	"testing"
)

func TestCreateGiftCodeService(t *testing.T) {
	config.InitClient()
	giftContent := structInfo2.GiftContentList{
		GoldCoins: 111,
		Diamonds:  222,
		Props:     333,
		Heroes:    444,
		Creeps:    555,
	}
	GiftCodeInfo := structInfo2.GiftCodeInfo{
		GiftDes:        "desc",
		AvailableTimes: 100000,
		ValidPeriod:    4,
		User:           "tom",
		ContentList:    giftContent,
	}
	code, e := CreateGiftCodeService(GiftCodeInfo)
	t.Log(code, e)
}

func TestGetGiftCodeInfoService(t *testing.T) {
	config.InitClient()
	GiftInfo, _ := GetGiftCodeInfoService("JI310XOC")
	t.Log(GiftInfo)
}

func TestVerifyFiftCodeService(t *testing.T) {
	config.InitClient()
	config.MongoClient()
	ContentInfo, err := VerifyFiftCodeService("KTWNWP1D", "H7SKR2")
	Reward := response.GeneralReward{}
	proto.Unmarshal(ContentInfo, &Reward)
	fmt.Println(Reward)
	t.Log(ContentInfo, err)
}
