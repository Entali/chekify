package env

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func GetTelegramID() int64 {
	userIDStr := os.Getenv("TELEGRAM_ID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		log.Fatalf("Invalid TELEGRAM_ID: %v", err)
	}
	return userID
}
