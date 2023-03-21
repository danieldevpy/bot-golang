package controller

import (
	"fmt"

	"github.com/danieldevpy/bot-golang/app/models"
	"gorm.io/gorm"
)

func GetUser(db *gorm.DB, number string, botid int) *models.Profile {
	var profile *models.Profile

	if err := db.Where("number = ? AND bot_id = ?", number, botid).First(&profile).Error; err != nil {
		var questionInicial int
		db.Table("bot_bot").Select("message_inicial_id").Where("id = ?", botid).Find(&questionInicial)
		profile = &models.Profile{Number: number, QuestionID: questionInicial, Activate: false, BotID: botid}
		if err := db.Create(&profile).Error; err != nil {
			fmt.Println("error crate")
		}
	}
	db.Preload("Question").Preload("Bot").First(&profile)
	return profile
}
