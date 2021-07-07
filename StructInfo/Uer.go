package StructInfo

type User struct{
	User string `json:"user"`
	UID string `json:"uid"`
	GoldCoins int `json:"gold_coins"`   //金币
	Diamonds int `json:"diamonds"`   //钻石
}