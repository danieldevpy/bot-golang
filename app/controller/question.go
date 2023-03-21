package controller

import (
	"fmt"

	"github.com/danieldevpy/bot-golang/app/models"
	"gorm.io/gorm"
)

func GetQuestionById(db *gorm.DB, idQuestion int) *models.Question {
	var question *models.Question
	if err := db.Where("id = ?", idQuestion).First(&question).Error; err != nil {
		fmt.Println("error")
	}
	return question
}

func GetParents(db *gorm.DB, question *models.Question) []models.Question {
	var questions []models.Question
	if err := db.Where("parent_question_id = ?", question.ID).Preload("ParentApp").Preload("ParentQuestion").Preload("OutPutQuestion").Order("id").Find(&questions).Error; err != nil {
		fmt.Println("error")
	}
	return questions
}

func ProcessQuestion(db *gorm.DB, profile *models.Profile, answer string) string {
	ProcessApp(db, profile)
	var questions_answers []models.Question

	if profile.Question.App {
		questions_answers = GetParents(db, profile.Question)
	} else {
		questions_answers = append(questions_answers, *profile.Question)
		parents := GetParents(db, profile.Question)
		questions_answers = append(questions_answers, parents...)
		fmt.Println("Não tem app aqui: ", profile.Question)
	}

	var lrange = 1
	object := ""
	for loop := range questions_answers {
		message := questions_answers[loop].Message
		if message[:1] != "!" {
			if questions_answers[loop].Index {
				object = fmt.Sprintf("%s*%d* - %s", object, lrange, message+"|")
				lrange = lrange + 1
			} else {
				object = object + message + "|"
			}
		}
	}
	fmt.Println(object)
	return object
}
