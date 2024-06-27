package repositories

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/models"
	"gorm.io/gorm"
)

func (r *repo) check(userName string, password string) *models.User {
	var user models.User

	if err := r.db.Where("UserName = ? and Password = ?", userName, password).First(&user).Error; err != nil {
		return nil
	}

	return &user
}

type repo struct {
	cfg *config.Config
	db  *gorm.DB
}

func newRepo() *repo {
	return &repo{
		cfg: config.GetInstance(),
		db:  database.GetInstance(),
	}
}
