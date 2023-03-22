package controller

import (
	"fmt"
	"regexp"

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

func GetAppSave(attribute string, db *gorm.DB, profile *models.Profile) string {
	var attributeSearch *models.Attribute
	var attributeSave *models.Save
	if err := db.Where("name = ? AND bot_id = ?", attribute, profile.BotID).First(&attributeSearch).Error; err != nil {
		fmt.Println(err)
	}
	if err := db.Where("attribute.id = ? AND profile_id = ? AND bot_id = ? ", attributeSearch.ID, profile.ID, profile.BotID).First(&attributeSave).Error; err != nil {
		fmt.Println(err)
	}
	return attributeSave.Armz
}

func ReplaceMatches(str string, re *regexp.Regexp, replacer func(string, *gorm.DB, *models.Profile) string, db *gorm.DB, profile *models.Profile) string {
	return re.ReplaceAllStringFunc(str, func(match string) string {
		substr := re.FindStringSubmatch(match)[1]
		return replacer(match[:1]+substr+match[len(match)-1:], db, profile)
	})
}
