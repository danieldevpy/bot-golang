package controller

import (
	"fmt"

	"github.com/danieldevpy/bot-golang/app/models"
	"gorm.io/gorm"
)

func ProcessApp(db *gorm.DB, profile *models.Profile) {
	if profile.Question.App {
		CheckingApp(db, profile, profile.Question)
	} else {
		parents := GetParents(db, profile.Question)

		for loop := range parents {
			if parents[loop].App {
				CheckingApp(db, profile, &parents[loop])
			}
		}
	}

}
func CheckingApp(db *gorm.DB, profile *models.Profile, question *models.Question) {
	db.Preload("Application").Find(question.ParentApp)
	switch question.ParentApp.Application.Action {
	case 1:
		fmt.Println("infos: ", question.ParentApp.AttributeID, profile.ID, profile.BotID)
		get_save := GetAttributeSave(db, question.ParentApp.AttributeID, profile.ID, profile.BotID)
		if get_save.ID > 0 {
			profile.Question = question.OutPutQuestion
			db.Preload("ParentApp").Preload("ParentQuestion").Preload("OutPutQuestion").Find(profile.Question)
			fmt.Println("Q: ", profile.Question)
			ProcessApp(db, profile)
		}
	}
}
