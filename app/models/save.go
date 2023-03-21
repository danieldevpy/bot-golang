package models

type Save struct {
	ID          int        `gorm:"primaryKey;column:id"`
	Attribute   *Attribute `gorm:"foreignKey:AttributeID"`
	AttributeID int        `gorm:"column:attribute_id"`
	Armz        string     `gorm:"column:armz"`
	Profile     *Profile   `gorm:"foreignKey:ProfileID"`
	ProfileID   int        `gorm:"column:profile_id"`
	Bot         *Bot
	BotID       int `gorm:"column:bot_id"`
}

func (s *Save) TableName() string {
	return "save_save"
}
