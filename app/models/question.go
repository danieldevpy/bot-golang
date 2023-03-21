package models

type Question struct {
	ID               int       `gorm:"primaryKey;column:id"`
	Message          string    `gorm:"column:message"`
	Index            bool      `gorm:"column:index"`
	Final            bool      `gorm:"column:final"`
	App              bool      `gorm:"column:app"`
	ParentApp        *CloneApp `gorm:"foreignKey:ParentAppID"`
	ParentAppID      int       `gorm:"column:parent_app_id"`
	ParentQuestion   *Question `gorm:"foreignKey:ParentQuestionID"`
	ParentQuestionID int       `gorm:"column:parent_question_id"`
	OutPutQuestion   *Question `gorm:"foreignKey:OutPutQuestionID"`
	OutPutQuestionID int       `gorm:"column:output_question_id"`
	Bot              *Bot      `gorm:"foreignKey:BotID"`
	BotID            int       `gorm:"column:bot_id"`
}

func (q *Question) TableName() string {
	return "question_question"
}
