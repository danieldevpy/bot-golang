package models

type Application struct {
	ID     int    `gorm:"primaryKey;column:id"`
	Name   string `gorm:"column:name"`
	Action int    `gorm:"column:action"`
}

func (a *Application) TableName() string {
	return "application_application"
}
