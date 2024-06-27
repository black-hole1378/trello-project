package handlers

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type workSpaceHandler struct {
	repository repositories.WorkSpaceRepoImpl
}

func NewWorkSpaceHandler() HandlerImpl {
	return &workSpaceHandler{
		repository: repositories.NewWorkSpaceRepo(),
	}
}

func (w workSpaceHandler) Create(c echo.Context) error {
	var newWorkSpace models.WorkSpace

	if err := c.Bind(&newWorkSpace); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	workSpace, err := w.repository.CreateWorkSpace(&newWorkSpace)

	if err != nil {
		return handleDBError(err, c)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"workSpace": workSpace,
	})

}

func (w workSpaceHandler) Delete(c echo.Context) error {
	id := c.Param("workspaceId")

	workSpaceId, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid workSpaceID!",
		})
	}

	if err := w.repository.DeleteWorkSpace(uint(workSpaceId)); err != nil {
		return handleDBError(err, c)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Successfully Deleted",
	})
}

func (w workSpaceHandler) Update(c echo.Context) error {
	id := c.Param("workspaceId")

	workSpaceId, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid workSpaceID!",
		})
	}

	var workSpace models.WorkSpace

	if err := c.Bind(&workSpace); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	workSpace.ID = uint(workSpaceId)

	if err := w.repository.UpdateWorkSpace(&workSpace); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Successfully updated!",
	})
}

func (w workSpaceHandler) GetAll(c echo.Context) error {
	workSpaces, err := w.repository.GetAllWorkSpaces()

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"workSpaces": workSpaces,
	})
}

func (w workSpaceHandler) Get(c echo.Context) error {
	id := c.Param("workspaceId")

	workSpaceId, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid workSpaceID!",
		})
	}

	workSpace, err := w.repository.GetWorkSpace(uint(workSpaceId))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"workSpace": workSpace,
	})
}
