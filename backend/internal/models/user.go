package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username      string          `gorm:"type:varchar;unique;not null"`
	Email         string          `gorm:"type:varchar;unique;not null"`
	Password      string          `gorm:"type:varchar;"`
	subTask       []SubTask       `gorm:"foreignKey:AssignedID;constraint:OnDelete:CASCADE;"`
	task          []Task          `gorm:"foreignKey:AssignedID;constraint:OnDelete:CASCADE;"`
	userWorkSpace []UserWorkSpace `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}
