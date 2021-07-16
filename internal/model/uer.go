package model

type User struct {
	UID       string `json:"uid"`
	GoldCoins int    `json:"gold_coins"` //金币
	Diamonds  int    `json:"diamonds"`   //钻石
}
