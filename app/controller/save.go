package controller

import (
	"github.com/danieldevpy/bot-golang/app/models"
	"gorm.io/gorm"
)

func GetAttributeSave(db *gorm.DB, attributeID int, profileID int, botID int) *models.Save {
	var save *models.Save
	db.Where("attribute_id = ? AND profile_id = ? AND bot_id = ?", attributeID, profileID, botID).First(&save)
	return save
}
