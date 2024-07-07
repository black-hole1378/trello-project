package models

import "gorm.io/gorm"

type UserWorkSpace struct {
	gorm.Model
	UserID      uint      `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	WorkSpaceID uint      `gorm:"not null"`
	WorkSpace   WorkSpace `gorm:"foreignKey:WorkSpaceID;constraint:OnDelete:CASCADE"`
	Role        string    `gorm:"type:role_states;default:'Standard User'"`
}
