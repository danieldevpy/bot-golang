package controller

import (
	"github.com/danieldevpy/bot-golang/app/models"
	"gorm.io/gorm"
)

func ProcessApp(db *gorm.DB, profile *models.Profile) {
	if profile.Question.App {
		CheckingApp(db, profile)
	} else {
		parents := GetParents(db, profile.Question)

		for loop := range parents {
			if parents[loop].App {
				profile.Question = &parents[loop]
				CheckingApp(db, profile)
			}
		}
	}

}
func CheckingApp(db *gorm.DB, profile *models.Profile) {
	db.Preload("Application").Find(profile.Question.ParentApp)
	switch profile.Question.ParentApp.Application.Action {
	case 1:
		get_save := GetAttributeSave(db, profile.Question.ParentApp.AttributeID, profile.ID, profile.BotID)
		if get_save.ID > 0 {
			profile.Question = profile.Question.OutPutQuestion
			db.Preload("ParentApp").Preload("ParentQuestion").Preload("OutPutQuestion").Find(profile.Question)
			ProcessApp(db, profile)
		}
	}
}

func ExecuteApp(db *gorm.DB, profile *models.Profile, answer string) {
	db.Preload("ParentApp").Preload("OutPutQuestion").Find(profile.Question)
	db.Preload("Application").Find(profile.Question.ParentApp)
	switch profile.Question.ParentApp.Application.Action {
	case 1:
		SaveAttribute(db, answer, profile)
		profile.Question = profile.Question.OutPutQuestion
	}
}
