package repositories

import (
	"backend/internal/models"
)

type CommentRepoImpl interface {
	Check(userName string, password string) (*models.User, error)
	DeleteComment(commentID uint, userID uint) error
	CreateComment(comment models.Comment) (*models.Comment, error)
	GetComment(commentID uint, userID uint) (*models.Comment, error)
	GetComments(taskID uint) ([]models.Comment, error)
}

type commentRepo struct {
	repo *repo
}

func (c *commentRepo) Check(userName string, password string) (*models.User, error) {
	var user models.User
	err := c.repo.db.Where("username = ? AND password = ?", userName, password).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *commentRepo) CreateComment(comment models.Comment) (*models.Comment, error) {
	err := c.repo.db.Create(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (c *commentRepo) DeleteComment(commentID uint, userID uint) error {
	var comment models.Comment
	err := c.repo.db.Where("id = ? AND user_id = ?", commentID, userID).First(&comment).Error
	if err != nil {
		return err
	}
	return c.repo.db.Delete(&comment).Error
}

func (c *commentRepo) GetComment(commentID uint, userID uint) (*models.Comment, error) {
	var comment models.Comment
	err := c.repo.db.Where("id = ? AND user_id = ?", commentID, userID).First(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (c *commentRepo) GetComments(taskID uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := c.repo.db.Where("task_id = ?", taskID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func NewCommentRepo() CommentRepoImpl {
	return &commentRepo{
		repo: newRepo(),
	}
}
