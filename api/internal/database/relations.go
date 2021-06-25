package database

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shujanpannag/iot_project_api/internal/model"
	"github.com/shujanpannag/iot_project_api/internal/util"
	"github.com/shujanpannag/iot_project_api/pkg/smtp"
	"gorm.io/gorm"
)

func GetNodelog(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		date := c.Param("date")
		var rv []model.Camlog
		q_date := date + "%"
		result := db.Where("datetime LIKE ?", q_date).Find(&rv)

		if result.Error != nil {
			c.JSON(409, gin.H{"error": result.Error})
			return
		}
		c.PureJSON(200, rv)
	}
}

func PostNodelog(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var f model.Camlog

		c.BindJSON(&f)

		result := db.Create(&f)

		if result.Error != nil {
			c.JSON(409, gin.H{"error": result.Error, "insert_query": f})
			return
		}

		// email notification
		subject := "GodsEye Logger\n"
		body := fmt.Sprintf("Someone at your doorstep\r\nName: %v\r\nTime: %v\r\n", f.RelationshipName, f.Datetime)
		to := util.GetEmailList(db, "Family")
		err := smtp.SendEmail(to, subject, body)

		c.JSON(201, gin.H{"db_error": result.Error, "insert_query": f, "rows_affected": result.RowsAffected, "smtp_error": err})
	}
}
