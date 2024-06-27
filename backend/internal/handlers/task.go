package handlers

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"net/http"
	"strconv"

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

	if err != nil {
		return handleDBError(err, c)
	}

	return c.JSON(http.StatusFound, echo.Map{
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

	var task models.Task

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	task.WorkSpaceID = uint(workSpaceId)

	user := t.repository.Check(c.Get("userName").(string), c.Get("password").(string))

	task.AssignedID = user.ID

	newTask, err := t.repository.CreateTask(task)

	if err != nil {
		return handleDBError(err, c)
	}

	return c.JSON(http.StatusCreated, echo.Map{
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
