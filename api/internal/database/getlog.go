package database

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shujanpannag/iot_project_api/internal/constant"
)

func GetLogReport(c *gin.Context) {
	currentTime := time.Now()
	fileName := fmt.Sprintf("%v_%v.%v", "geserver", currentTime.Format("200601021504"), "log")

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")

	c.FileAttachment(constant.SERVER_LOG_PATH, fileName)
}
