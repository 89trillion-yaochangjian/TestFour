package handler

import (
	"MongoGift/StructInfo"
	"MongoGift/internal/service"
	"fmt"
)

//管理后台调用 - 创建礼品码

func CreateGiftCodeHandler(giftCodeInfo StructInfo.GiftCodeInfo) (StructInfo.MesInfo,error){
	codeRes,err:= service.CreateGiftCodeService(giftCodeInfo)
	if err != nil {
		return StructInfo.MesInfo{Msg: "创建礼包码失败",ER: err},nil
	}
	return StructInfo.MesInfo{Msg: "创建礼包码成功",Data: codeRes},nil
}


//管理后台调用 - 查询礼品码信息

func GetFiftCodeInfoHandler(code string) (StructInfo.MesInfo,error){
	giftCodeInfo,e := service.GetGiftCodeInfoService(code)
	if e!=nil {
		fmt.Println("GetFiftCodeInfoHandler")
		return StructInfo.MesInfo{Msg: "查询礼品码信息失败",ER: e},nil
	}
	return StructInfo.MesInfo{Msg: "查询礼品码信息成功",Data: giftCodeInfo},nil
}


//客户端调用 - 验证礼品码

func VerifyFiftCodeHandler(code string,user string) ([]byte,error)  {
	giftCodeInfo, er := service.VerifyFiftCodeService(code,user)
	return giftCodeInfo,er
}