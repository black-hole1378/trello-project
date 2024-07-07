package handlers

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type taskHandler struct {
	repository repositories.TaskRepoImpl
}

func NewTaskHandler() HandlerImpl {
	return &taskHandler{
		repository: repositories.NewTaskRepository(),
	}
}

func (t *taskHandler) Update(c echo.Context) error {
	workSpaceID, taskID, err := getTaskID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	var task models.Task

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	task.ID = taskID
	task.WorkSpaceID = workSpaceID

	if err := t.repository.UpdateTask(task); err != nil {
		return handleDBError(err, c)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Successfully Updated the task!",
	})
}

func (t *taskHandler) GetAll(c echo.Context) error {
	id := c.Param("workspaceId")
	workSpaceId, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid Workspace!",
		})
	}

	tasks, err := t.repository.GetTasks(uint(workSpaceId))
	fmt.Println(tasks)
	if err != nil {
		return handleDBError(err, c)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"tasks": tasks,
	})

}

func getTaskID(c echo.Context) (uint, uint, error) {
	id := c.Param("workspaceId")

	tID := c.Param("taskID")

	workSpaceId, _ := strconv.ParseUint(id, 10, 32)

	taskID, err := strconv.ParseUint(tID, 10, 32)

	return uint(workSpaceId), uint(taskID), err
}

func (t *taskHandler) Get(c echo.Context) error {
	workSpaceID, taskID, err := getTaskID(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid Workspace or Task ID!",
		})
	}

	task, err := t.repository.GetTask(taskID, workSpaceID)

	if err != nil {
		return handleDBError(err, c)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"task": task,
	})
}

func (t *taskHandler) Create(c echo.Context) error {
	id := c.Param("workspaceId")

	workSpaceId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid Workspace ID!",
		})
	}

	var request struct {
		Title         string `json:"title"`
		Description   string `json:"description"`
		EstimatedTime string `json:"estimatedTime"` // Assuming frontend sends this as a string
		DueDate       string `json:"dueDate"`
		ImageUrl      string `json:"imageUrl"`
		Priority      string `json:"priority"`
		Status        string `json:"status"`
		Assigned      string `json:"userName"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request body!",
		})
	}

	// Parsing due date
	dueDate, err := time.Parse("2006-01-02", request.DueDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid Due Date format! Use YYYY-MM-DD.",
		})
	}

	// Parsing estimated time (assuming it is in hours and converting to Duration)
	estimatedTime, err := time.ParseDuration(request.EstimatedTime + "h")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid Estimated Time format!",
		})
	}

	userRepository := repositories.NewUserRepository()
	user, err := userRepository.GetUserByUserName(request.Assigned)
	if err != nil {
		return handleDBError(err, c)
	}

	// Assigning the request data to the Task model
	task := models.Task{
		Title:         request.Title,
		Description:   request.Description,
		EstimatedTime: estimatedTime,
		DueDate:       dueDate,
		ImageUrl:      request.ImageUrl,
		Priority:      request.Priority,
		Status:        request.Status,
		AssignedID:    uint(user.ID),
		WorkSpaceID:   uint(workSpaceId),
	}

	// Create the task in the database
	newTask, err := t.repository.CreateTask(task)
	if err != nil {
		fmt.Println("Error creating task:", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error creating task!",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"task": newTask,
	})
}

func (t *taskHandler) Delete(c echo.Context) error {
	workSpaceID, taskID, err := getTaskID(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid Workspace or Task ID!",
		})
	}

	if err := t.repository.DeleteTask(taskID, workSpaceID); err != nil {
		return handleDBError(err, c)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Successfully deleted the Task!",
	})

}
