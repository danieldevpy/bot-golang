package models

type Bot struct {
	ID               int       `gorm:"primaryKey;column:id"`
	Name             string    `gorm:"column:name"`
	MessageInicial   *Question `gorm:"foreignKey:MessageInicialID"`
	MessageInicialID int       `gorm:"column:message_inicial_id"`
	Activate         bool      `gorm:"column:activate"`
}

func (b *Bot) TableName() string {
	return "bot_bot"
}
