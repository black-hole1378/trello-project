package repositories

import (
	"backend/internal/models"
)

type UserWorkSpaceRepoImpl interface {
	Check(userName string, password string) *models.User
	DeleteUserWorkSpace(workSpaceID uint, userID uint) error
	CreateUserWorkSpace(workSpace models.UserWorkSpace) (*models.UserWorkSpace, error)
	UpdateUserWorkSpace(workSpace models.UserWorkSpace) error
	GetUsersWorkSpace(workSpaceID uint) ([]models.User, error)
}

type userWorkSpaceRepo struct {
	*repo
}

// Check implements UserWorkSpaceRepoImpl.
func (u *userWorkSpaceRepo) Check(userName string, password string) *models.User {
	return u.check(userName, password)
}

// CreateUserWorkSpace implements UserWorkSpaceRepoImpl.
func (u *userWorkSpaceRepo) CreateUserWorkSpace(workSpace models.UserWorkSpace) (*models.UserWorkSpace, error) {
	if err := u.db.Create(&workSpace).Error; err != nil {
		return nil, err
	}

	return &workSpace, nil
}

// DeleteUserWorkSpace implements UserWorkSpaceRepoImpl.
func (u *userWorkSpaceRepo) DeleteUserWorkSpace(workSpaceID uint, userID uint) error {
	return u.db.Where("UserID = ?", userID).Delete(models.UserWorkSpace{}, workSpaceID).Error
}

// GetUsersWorkSpace implements UserWorkSpaceRepoImpl.
func (u *userWorkSpaceRepo) GetUsersWorkSpace(workSpaceID uint) ([]models.User, error) {
	var users []models.User

	if err := u.db.Joins("JOIN user_work_spaces uws ON uws.user_id = users.id").
		Where("uws.work_space_id = ?", workSpaceID).
		Preload("UserWorkSpaces").
		Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUserWorkSpace implements UserWorkSpaceRepoImpl.
func (u *userWorkSpaceRepo) UpdateUserWorkSpace(workSpace models.UserWorkSpace) error {
	if err := u.db.Save(&workSpace).Error; err != nil {
		return err
	}
	return nil
}

func NewUserWorkSpace() UserWorkSpaceRepoImpl {
	return &userWorkSpaceRepo{
		repo: newRepo(),
	}
}
