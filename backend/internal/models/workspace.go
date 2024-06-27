package models

import "gorm.io/gorm"

type WorkSpace struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:varchar"`
	task        []Task `gorm:"foreignKey:WorkSpaceID;constraint:OnDelete:CASCADE;"`
}
