package controller

import (
	"fmt"

	"github.com/danieldevpy/bot-golang/app/models"
	"gorm.io/gorm"
)

func GetResponse(db *gorm.DB, botid int, number string, answer string) (*models.Profile, string) {

	profile := GetUser(db, number, botid)
	fmt.Println("inicio: ", profile)

	if answer == "0" {
		profile.Question = GetQuestionById(db, profile.Bot.MessageInicialID)
	}

	bot_answer := ProcessQuestion(db, profile, answer)

	return profile, bot_answer
}
