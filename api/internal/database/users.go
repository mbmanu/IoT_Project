package database

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/shujanpannag/iot_project_api/internal/model"
	"github.com/shujanpannag/iot_project_api/internal/util"
	"github.com/shujanpannag/iot_project_api/pkg/smtp"

	"gorm.io/gorm"
)

func GetUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		rel := c.Param("rel")
		var rv model.Relationship
		db.Where("name = ?", rel).Find(&rv)

		c.PureJSON(200, rv)
	}
}

func PutUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		name := c.Param("name")

		var f model.Relationship

		c.BindJSON(&f)

		if f.Name != name {
			c.JSON(400, gin.H{"error": true})
			return
		}

		// email notification
		subject := "GodsEye Logger\n"
		body := fmt.Sprintf("User Relationship Updated\r\nRelation Name: %v\r\nRelation Type: %v\r\nAt: %v\r\n", f.Name, f.Rel, time.Now().Format("02 Jan 06 15:04"))
		to := util.GetEmailList(db, "Family")
		err := smtp.SendEmail(to, subject, body)

		result := db.Model(&f).Updates(model.Relationship{Name: f.Name, Rel: f.Rel, Email: f.Email})
		c.JSON(200, gin.H{"db_error": result.Error, "update_query": f, "rows_affected": result.RowsAffected, "smtp_error": err})
	}
}

func PostUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var f model.Relationship

		c.BindJSON(&f)

		result := db.Create(&f)
		if result.Error != nil {
			c.JSON(400, gin.H{"db_error": result.Error, "update_query": f, "rows_affected": result.RowsAffected})
			return
		}

		// email notification
		subject := "GodsEye Logger\n"
		body := fmt.Sprintf("New User Relationship Added\r\nRelation Name: %v\r\nRelation Type: %v\r\nAt: %v\r\n", f.Name, f.Rel, time.Now().Format("02 Jan 06 15:04"))
		to := util.GetEmailList(db, "Family")
		err := smtp.SendEmail(to, subject, body)

		c.JSON(201, gin.H{"db_error": result.Error, "update_query": f, "rows_affected": result.RowsAffected, "smtp_error": err})
	}
}

func DelUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		name := c.Param("name")

		f := model.Relationship{Name: name}

		result := db.Delete(&f)
		c.JSON(200, gin.H{"db_error": result.Error, "rows_affected": result.RowsAffected})
	}
}
