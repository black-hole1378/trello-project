package repositories

import (
	"backend/internal/models"
	"backend/internal/utils"
)

type UserRepoImpl interface {
	DeleteUser(userID uint, userName string) error
	CreateUser(user *models.User) (models.User, error)
	UpdateUser(updateUser *models.User, userName string, userID uint) error
	GetUser(userID uint, userName string) (*models.User, error)
	Check(userName string, password string) *models.User
	GetAllUsers() ([]models.User, error)
	GetUserByUserName(userName string) (*models.User, error)
}

type userRepo struct {
	*repo
}

// GetUserByUserName implements UserRepoImpl.
func (r *userRepo) GetUserByUserName(userName string) (*models.User, error) {
	var user *models.User

	if err := r.db.Where("Username = ?", userName).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserRepository() UserRepoImpl {
	return &userRepo{
		repo: newRepo(),
	}
}

func (r *userRepo) DeleteUser(userID uint, userName string) error {
	return r.db.Where("UserName = ?", userName).Delete(&models.User{}, userID).Error
}

func (r *userRepo) CreateUser(user *models.User) (models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return *user, nil
}

func (r *userRepo) UpdateUser(updateUser *models.User, userName string, userID uint) error {
	user, err := r.GetUser(userID, userName)
	if err != nil {
		return err
	}

	switch {
	case updateUser.Email != "":
		user.Email = updateUser.Email
	case updateUser.Username != "":
		user.Username = updateUser.Username
	case updateUser.Password != "":
		user.Password = utils.HashedPassword(updateUser.Password)
	}

	if err := r.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepo) GetUser(userID uint, userName string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("UserName = ?", userName).First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) Check(userName string, password string) *models.User {
	return r.check(userName, password)
}
