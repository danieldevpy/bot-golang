package models

type Attribute struct {
	ID    int    `gorm:"primaryKey;column:id"`
	Name  string `gorm:"column:name"`
	Bot   *Bot
	BotID int `gorm:"column:bot_id"`
}

func (at *Attribute) TableName() string {
	return "attribute_attribute"
}
