package models

import "gorm.io/gorm"

type WorkSpace struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:varchar"`
	Task        []Task `gorm:"foreignKey:WorkSpaceID;constraint:OnDelete:CASCADE;"`
}
