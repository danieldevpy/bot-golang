package models

type Profile struct {
	ID         int    `gorm:"primaryKey;column:id"`
	Number     string `gorm:"column:number"`
	Activate   bool   `gorm:"column:activate"`
	Question   *Question
	QuestionID int `gorm:"column:question_id"`
	Bot        *Bot
	BotID      int `gorm:"column:bot_id"`
}

func (p *Profile) TableName() string {
	return "profille_profile"
}
