package models

import "gorm.io/gorm"

type SubTask struct {
	gorm.Model
	TaskID      uint   `gorm:"foreignKey:ID;not null;"`
	Task        Task   `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE;"`
	Title       string `gorm:"type:varchar;not null"`
	IsCompleted string `gorm:"type:completed_state;default:'NO'"`
	AssignedID  uint   `gorm:"foreignKey:ID;not null;"`
	AssignedTo  User   `gorm:"foreignKey:AssignedID;constraint:OnDelete:CASCADE;"`
}
