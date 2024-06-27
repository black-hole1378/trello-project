package repositories

import (
	"backend/internal/models"
)

type WorkSpaceRepoImpl interface {
	DeleteWorkSpace(workSpaceID uint) error
	CreateWorkSpace(workSpace *models.WorkSpace) (models.WorkSpace, error)
	UpdateWorkSpace(updatedWorkSpace *models.WorkSpace) error
	GetAllWorkSpaces() ([]models.WorkSpace, error)
	GetWorkSpace(workSpaceID uint) (models.WorkSpace, error)
}

type workSpaceRepo struct {
	*repo
}

func NewWorkSpaceRepo() WorkSpaceRepoImpl {
	return &workSpaceRepo{
		repo: newRepo(),
	}
}

func (r *workSpaceRepo) DeleteWorkSpace(workSpaceID uint) error {
	return r.db.Delete(&models.WorkSpace{}, workSpaceID).Error
}

func (r *workSpaceRepo) CreateWorkSpace(workSpace *models.WorkSpace) (models.WorkSpace, error) {
	if err := r.db.Create(&workSpace).Error; err != nil {
		return models.WorkSpace{}, err
	}
	return *workSpace, nil
}

func (r *workSpaceRepo) UpdateWorkSpace(updatedWorkSpace *models.WorkSpace) error {
	workSpace, err := r.GetWorkSpace(updatedWorkSpace.ID)
	if err != nil {
		return err
	}

	switch {
	case updatedWorkSpace.Name != "":
		workSpace.Name = updatedWorkSpace.Name
	case updatedWorkSpace.Description != "":
		workSpace.Description = updatedWorkSpace.Description

	}

	if err := r.db.Save(&workSpace).Error; err != nil {
		return err
	}

	return nil
}

func (r *workSpaceRepo) GetAllWorkSpaces() ([]models.WorkSpace, error) {
	var workSpaces []models.WorkSpace
	if err := r.db.Find(&workSpaces).Error; err != nil {
		return nil, err
	}
	return workSpaces, nil
}

func (r *workSpaceRepo) GetWorkSpace(workSpaceID uint) (models.WorkSpace, error) {
	var workSpace models.WorkSpace
	if err := r.db.First(&workSpace, workSpaceID).Error; err != nil {
		return models.WorkSpace{}, err
	}
	return workSpace, nil
}
