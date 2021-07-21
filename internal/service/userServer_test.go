package service

import (
	"MongoGift/internal/config"
	"testing"
)

func TestUserLoginServer(t *testing.T) {
	config.InitClient()
	config.MongoClient()
	ContentInfo, err := UserLoginServer("tom")
	t.Log(ContentInfo, err)
}
