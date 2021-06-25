package util

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shujanpannag/iot_project_api/internal/constant"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

func GinLogger() (*os.File, gin.HandlerFunc) {
	f, err := os.OpenFile(constant.SERVER_LOG_PATH, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log := logrus.New()
	log.SetOutput(f)

	return f, ginlogrus.Logger(log)
}
