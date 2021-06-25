package util

import (
	"github.com/shujanpannag/iot_project_api/internal/model"
	"gorm.io/gorm"
)

func GetEmailList(db *gorm.DB, relType string) []string {
	rv := []model.Relationship{}

	db.Where("rel = ?", relType).Find(&rv)
	var to []string
	for _, x := range rv {
		to = append(to, x.Email)
	}

	return to
}
