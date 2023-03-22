package controller

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/danieldevpy/bot-golang/app/models"
	"gorm.io/gorm"
)

func GetResponse(db *gorm.DB, botid int, number string, answer string) (*models.Profile, string) {
	profile := GetUser(db, number, botid)

	if answer == "0" {
		profile.Question = GetQuestionById(db, profile.Bot.MessageInicialID)
	}

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
		if !profile.Question.Index {
			questions_answers = append(questions_answers, *profile.Question)
		}
		parents := GetParents(db, profile.Question)
		questions_answers = append(questions_answers, parents...)

		lloop := 1
		for loop := range questions_answers {
			if questions_answers[loop].Index {
				if answer == strconv.Itoa(lloop) {
					profile.Question = &questions_answers[loop]
					questions_answers = GetParents(db, profile.Question)
					break
				}
				lloop = lloop + 1
			}
		}

	}
	if profile.Question.Final {
		questions_answers = append(questions_answers, GetFinal(db, profile))
		profile.Question = GetQuestionById(db, profile.Bot.MessageInicialID)
	}

	var lrange = 1
	object := ""
	for loop := range questions_answers {
		message := questions_answers[loop].Message
		if message[:1] != "!" {
			re := regexp.MustCompile(`\{([^}]+)\}`)
			matches := re.FindAllStringSubmatch(questions_answers[loop].Message, -1)
			if len(matches) > 0 {
				message = ReplaceMatches(message, re, GetAppSave, db, profile)
			}
			if questions_answers[loop].Index {
				object = fmt.Sprintf("%s*%d* - %s", object, lrange, message+"|")
				lrange = lrange + 1
			} else {
				object = object + message + "|"
			}
		}
	}

	return profile, object
}
