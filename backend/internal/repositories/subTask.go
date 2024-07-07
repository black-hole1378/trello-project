package repositories

import "backend/internal/models"

type SubTaskRepoImpl interface {
	Check(userName string, password string) *models.User
	DeleteSubTask(taskID uint, subTaskID uint) error
	CreateSubTask(task models.SubTask) (*models.SubTask, error)
	UpdateSubTask(task models.SubTask) error
	GetSubTask(taskID uint, workSpaceID uint) (*models.SubTask, error)
	GetSubTasks(subTaskID uint) ([]models.SubTask, error)
}

type subTaskRepo struct {
	*repo
}

// Check implements SubTaskRepoImpl.
func (s *subTaskRepo) Check(userName string, password string) *models.User {
	return s.check(userName, password)
}

// CreateSubTask implements SubTaskRepoImpl.
func (s *subTaskRepo) CreateSubTask(task models.SubTask) (*models.SubTask, error) {
	if err := s.db.Create(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

// DeleteSubTask implements SubTaskRepoImpl.
func (s *subTaskRepo) DeleteSubTask(taskID uint, subTaskID uint) error {
	return s.db.Where("TaskID = ?", taskID).Delete(models.SubTask{}, subTaskID).Error
}

// GetSubTask implements SubTaskRepoImpl.
func (s *subTaskRepo) GetSubTask(taskID uint, subTaskID uint) (*models.SubTask, error) {
	var task models.SubTask

	if err := s.db.Where("TaskID = ?", taskID).First(&task, subTaskID).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

// GetSubTasks implements SubTaskRepoImpl.
func (s *subTaskRepo) GetSubTasks(taskID uint) ([]models.SubTask, error) {
	var tasks []models.SubTask

	if err := s.db.Where("TaskID = ? ", taskID).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

// UpdateSubTask implements SubTaskRepoImpl.
func (s *subTaskRepo) UpdateSubTask(task models.SubTask) error {
	findedTask, err := s.GetSubTask(task.ID, task.TaskID)

	if err != nil {
		return err
	}

	switch {
	case task.Title != "":
		findedTask.Title = task.Title
	case task.IsCompleted != "":
		findedTask.IsCompleted = task.IsCompleted
	}

	if err := s.db.Save(&findedTask).Error; err != nil {
		return err
	}

	return nil
}

func NewSubTaskRepository() SubTaskRepoImpl {
	return &subTaskRepo{
		repo: newRepo(),
	}
}
