package repositories

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/models"
	"backend/internal/utils"

	"gorm.io/gorm"
)

func (r *repo) check(userName string, password string) *models.User {
	var user models.User

	if err := r.db.Where("UserName = ?", userName).First(&user).Error; err != nil {
		return nil
	}

	if !utils.ComparePasswordHash(password, user.Password) {
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
