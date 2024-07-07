package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username       string          `gorm:"type:varchar;unique;not null"`
	Email          string          `gorm:"type:varchar;unique;not null"`
	Password       string          `gorm:"type:varchar;"`
	SubTasks       []SubTask       `gorm:"foreignKey:AssignedID;constraint:OnDelete:CASCADE"`
	Tasks          []Task          `gorm:"foreignKey:AssignedID;constraint:OnDelete:CASCADE"`
	UserWorkSpaces []UserWorkSpace `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
