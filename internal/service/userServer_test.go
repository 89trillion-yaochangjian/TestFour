package service

import (
	"MongoGift/internal/utils"
	"testing"
)

func TestUserLoginServer(t *testing.T) {
	utils.InitClient()
	utils.MongoClient()
	ContentInfo, err := UserLoginServer("tom")
	t.Log(ContentInfo, err)
}
