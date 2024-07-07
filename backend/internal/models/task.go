package models

import (
	"time"
)

type Task struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	Title         string    `gorm:"type:varchar;not null"`
	Description   string    `gorm:"type:varchar"`
	ImageUrl      string    `gorm:"type:varchar"`
	WorkSpaceID   uint      `gorm:"not null"`
	Workspace     WorkSpace `gorm:"foreignKey:WorkSpaceID;constraint:OnDelete:CASCADE;"`
	Priority      string    `gorm:"type:varchar"`
	AssignedID    uint      `gorm:"not null"`
	AssignedTo    User      `gorm:"foreignKey:AssignedID;constraint:OnDelete:CASCADE;"`
	SubTask       []SubTask `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE;"`
	Status        string    `gorm:"type:task_status;default:'In Progress'"`
	EstimatedTime time.Duration
	ActualTime    time.Duration
	DueDate       time.Time
}
