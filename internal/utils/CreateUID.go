package utils

import (
	"github.com/gofrs/uuid"
	"log"
)

func CreateUID() string {
	u2, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	return u2.String()
}