package controller

import (
	"fmt"
	"strconv"

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
func ProcessQuestion(db *gorm.DB, profile *models.Profile, answer string) string {
	var questions_answers []models.Question

	if !profile.Activate {
		profile.Activate = true
		questions_answers = append(questions_answers, *profile.Question)
	}

	if profile.Question.App {
		ExecuteApp(db, profile, answer)
		questions_answers = append(questions_answers, *profile.Question)
		questions_answers = append(questions_answers, GetParents(db, profile.Question)...)
	} else {
		ProcessApp(db, profile)
		fmt.Println("chamado fora de app")
		if !profile.Question.Index {
			questions_answers = append(questions_answers, *profile.Question)
		}
		parents := GetParents(db, profile.Question)
		questions_answers = append(questions_answers, parents...)
		fmt.Println("escolhas disponivel: ", questions_answers)
		lloop := 1
		for loop := range questions_answers {
			if questions_answers[loop].Index {
				if answer == strconv.Itoa(lloop) {
					fmt.Println("caiu aqui escolha: ", questions_answers[loop])
					profile.Question = &questions_answers[loop]
					questions_answers = GetParents(db, profile.Question)
					break
				}
				lloop = lloop + 1
			}
		}

	}
	fmt.Println("Q:", profile.Question)
	if profile.Question.Final {
		questions_answers = append(questions_answers, GetFinal(db, profile))
		profile.Question = GetQuestionById(db, profile.Bot.MessageInicialID)
	}

	lrange := 1
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

	return object
}
