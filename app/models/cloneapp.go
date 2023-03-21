package models

type CloneApp struct {
	ID            int          `gorm:"primaryKey;column:id"`
	Application   *Application `gorm:"foreignKey:ApplicationID"`
	ApplicationID int          `gorm:"column:application_id"`
	Attribute     *Attribute   `gorm:"foreignKey:AttributeID"`
	AttributeID   int          `gorm:"column:attribute_id"`
	Bot           *Bot
	BotID         int `gorm:"column:bot_id"`
}

func (c *CloneApp) TableName() string {
	return "cloneapp_cloneapp"
}
