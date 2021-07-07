package service

import (
	"MongoGift/StructInfo"
	"MongoGift/internal/utils"
	"testing"
)

func TestCreateGiftCodeService(t *testing.T) {
	utils.InitClient()
	giftContent := StructInfo.GiftContentList{
		GoldCoins:111,
		Diamonds:222,
		Props:333,
		Heroes:444,
		Creeps:555,
	}
	GiftCodeInfo := StructInfo.GiftCodeInfo{
		GiftDes:"desc",
		AvailableTimes:100000,
		ValidPeriod:4,
		User: "tom",
		ContentList:giftContent,
	}
	code,e := CreateGiftCodeService(GiftCodeInfo)
	t.Log(code,e)
}

func TestGetGiftCodeInfoService(t *testing.T) {
	utils.InitClient()
	GiftInfo,_ := GetGiftCodeInfoService("JI310XOC")
	t.Log(GiftInfo)
}

//func TestVerifyFiftCodeService(t *testing.T) {
//	utils.InitClient()
//	ContentInfo := VerifyFiftCodeService("A4UJTDLV","tom")
//	t.Log(ContentInfo)
//}

func TestVerifyFiftCodeService(t *testing.T) {
	utils.InitClient()
	utils.MongoClient()
	ContentInfo,err := VerifyFiftCodeService("A4UJTDLV","tom")
	t.Log(ContentInfo,err)
}