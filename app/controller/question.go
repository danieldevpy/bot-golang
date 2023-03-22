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

func GetFinal(db *gorm.DB, profile *models.Profile) models.Question {
	var question *models.Question
	err := db.Where("id = ?", profile.Question.OutPutQuestionID).Find(&question).Error
	if err != nil {
		fmt.Println("error")
	}
	return *question
}
