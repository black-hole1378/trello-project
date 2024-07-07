package models

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	TaskID    uint      `gorm:"not null"`
	Task      Task      `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE;"`
	UserID    uint      `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
