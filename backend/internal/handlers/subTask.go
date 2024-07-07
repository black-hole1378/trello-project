package handlers

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type subTaskHandler struct {
	repository repositories.SubTaskRepoImpl
}

func NewSubTaskHandler() HandlerImpl {
	return &subTaskHandler{
		repository: repositories.NewSubTaskRepository(),
	}
}

// Create implements HandlerImpl.
func (s *subTaskHandler) Create(c echo.Context) error {
	id := c.Param("taskID")

	taskID, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid task ID!",
		})
	}

	var request struct {
		UserName   string `json:"userName"`
		Title      string `json:"title"`
		IsComplete string `json:"status"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	repository := repositories.NewUserRepository()

	user, err := repository.GetUserByUserName(request.UserName)

	if err != nil {
		return handleDBError(err, c)
	}

	subTask := models.SubTask{
		Title:       request.Title,
		AssignedID:  user.ID,
		TaskID:      uint(taskID),
		IsCompleted: request.IsComplete,
	}

	newTask, err := s.repository.CreateSubTask(subTask)

	if err != nil {
		return handleDBError(err, c)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"subTask": newTask,
	})
}

func getSubTaskID(c echo.Context) (uint, uint, error) {
	id := c.Param("taskID")

	tID := c.Param("subtaskID")

	taskID, _ := strconv.ParseUint(id, 10, 32)

	subTaskID, err := strconv.ParseUint(tID, 10, 32)

	return uint(taskID), uint(subTaskID), err
}

// Delete implements HandlerImpl.
func (s *subTaskHandler) Delete(c echo.Context) error {
	taskID, subTaskID, err := getSubTaskID(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid Workspace or Task ID!",
		})
	}

	if err := s.repository.DeleteSubTask(taskID, subTaskID); err != nil {
		return handleDBError(err, c)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Successfully deleted the SubTask!",
	})
}

// Get implements HandlerImpl.
func (s *subTaskHandler) Get(c echo.Context) error {
	taskID, subTaskID, err := getSubTaskID(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid subTask ID or Task ID!",
		})
	}

	task, err := s.repository.GetSubTask(taskID, subTaskID)

	if err != nil {
		return handleDBError(err, c)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"subTask": task,
	})
}

// GetAll implements HandlerImpl.
func (s *subTaskHandler) GetAll(c echo.Context) error {
	id := c.Param("taskID")
	taskID, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid Task ID!",
		})
	}

	tasks, err := s.repository.GetSubTasks(uint(taskID))

	if err != nil {
		return handleDBError(err, c)
	}

	return c.JSON(http.StatusFound, echo.Map{
		"subTasks": tasks,
	})
}

// Update implements HandlerImpl.
func (s *subTaskHandler) Update(c echo.Context) error {
	taskID, subTaskID, err := getSubTaskID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	var task models.SubTask

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	task.ID = subTaskID
	task.TaskID = taskID

	if err := s.repository.UpdateSubTask(task); err != nil {
		return handleDBError(err, c)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Successfully Updated the subTask!",
	})

}
