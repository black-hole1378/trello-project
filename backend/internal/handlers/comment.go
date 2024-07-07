package handlers

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	repository repositories.CommentRepoImpl
}

func (h *CommentHandler) Create(c echo.Context) error {
	id := c.Param("taskID")
	taskID, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		fmt.Println("error", err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid taskID",
		})
	}

	var request struct {
		UserID  string `json:"userID"`
		Content string `json:"content"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	userID, err := strconv.ParseUint(request.UserID, 10, 32)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid userID",
		})
	}

	comment := models.Comment{
		TaskID:  uint(taskID),
		UserID:  uint(userID),
		Content: request.Content,
	}

	newComment, err := h.repository.CreateComment(comment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, newComment)
}

func (h *CommentHandler) Delete(c echo.Context) error {
	commentID, err := strconv.ParseUint(c.Param("commentID"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid comment ID"})
	}

	userID, err := strconv.ParseUint(c.Param("userID"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID"})
	}

	err = h.repository.DeleteComment(uint(commentID), uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Comment deleted successfully"})
}

func (h *CommentHandler) Get(c echo.Context) error {
	commentID, err := strconv.ParseUint(c.Param("commentID"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid comment ID"})
	}

	userID, err := strconv.ParseUint(c.Param("userID"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID"})
	}

	comment, err := h.repository.GetComment(uint(commentID), uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, comment)
}

func (h *CommentHandler) GetAll(c echo.Context) error {
	taskID, err := strconv.ParseUint(c.Param("taskID"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid task ID"})
	}

	comments, err := h.repository.GetComments(uint(taskID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) Update(c echo.Context) error {
	var request struct {
		ID      uint   `json:"id" validate:"required"`
		Content string `json:"content" validate:"required"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	comment, err := h.repository.GetComment(request.ID, c.Get("userID").(uint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	comment.Content = request.Content

	updatedComment, err := h.repository.CreateComment(*comment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, updatedComment)
}

func NewCommentHandler() *CommentHandler {
	return &CommentHandler{
		repository: repositories.NewCommentRepo(),
	}
}
