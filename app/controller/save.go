package controller

import (
	"fmt"

	"github.com/danieldevpy/bot-golang/app/models"
	"gorm.io/gorm"
)

func GetAttributeSave(db *gorm.DB, attributeID int, profileID int, botID int) *models.Save {
	var save *models.Save
	db.Where("attribute_id = ? AND profile_id = ? AND bot_id = ?", attributeID, profileID, botID).First(&save)
	return save
}
func SaveAttribute(db *gorm.DB, armz string, profile *models.Profile) {
	save := models.Save{Armz: armz, AttributeID: profile.Question.ParentApp.AttributeID, ProfileID: profile.ID, BotID: profile.BotID}
	if err := db.Create(&save).Error; err != nil {
		fmt.Println("error crate")
	}
}
