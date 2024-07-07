package repositories

import (
	"backend/internal/models"
)

type TaskRepoImpl interface {
	Check(userName string, password string) *models.User
	DeleteTask(taskID uint, workSpaceID uint) error
	CreateTask(task models.Task) (*models.Task, error)
	UpdateTask(task models.Task) error
	GetTask(taskID uint, workSpaceID uint) (*models.Task, error)
	GetTasks(workSpaceID uint) ([]models.Task, error)
}

type taskRepo struct {
	*repo
}

func NewTaskRepository() TaskRepoImpl {
	return &taskRepo{
		repo: newRepo(),
	}
}

func (r *taskRepo) Check(userName string, password string) *models.User {
	return r.repo.check(userName, password)
}

func (r *taskRepo) DeleteTask(taskID uint, workSpaceID uint) error {
	return r.db.Where("WorkSpaceID = ?", workSpaceID).Delete(models.Task{}, taskID).Error
}

func (r *taskRepo) CreateTask(task models.Task) (*models.Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *taskRepo) UpdateTask(task models.Task) error {
	findedTask, err := r.GetTask(task.ID, task.WorkSpaceID)

	if err != nil {
		return err
	}

	switch {
	case task.Description != "":
		findedTask.Description = task.Description
	case task.Title != "":
		findedTask.Title = task.Title
	case task.ImageUrl != "":
		findedTask.ImageUrl = task.ImageUrl
	case task.ActualTime != 0:
		findedTask.ActualTime = task.ActualTime
	case task.EstimatedTime != 0:
		findedTask.EstimatedTime = task.EstimatedTime
	case task.Priority != "":
		findedTask.Priority = task.Priority
	case task.Status != "":
		findedTask.Status = task.Status
	}

	if err := r.db.Save(&findedTask).Error; err != nil {
		return err
	}

	return nil
}

func (r *taskRepo) GetTask(taskID uint, workSpaceID uint) (*models.Task, error) {
	var task models.Task

	if err := r.db.Where("WorkSpaceID = ?", workSpaceID).First(&task, taskID).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *taskRepo) GetTasks(workSpaceID uint) ([]models.Task, error) {
	var tasks []models.Task

	if err := r.db.Where("work_space_id = ? ", workSpaceID).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}
