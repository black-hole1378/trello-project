package handlers

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userWorkSpaceHandler struct {
	repository repositories.UserWorkSpaceRepoImpl
}

func NewUserWorkSpaceHandler() HandlerImpl {
	return &userWorkSpaceHandler{
		repository: repositories.NewUserWorkSpace(),
	}
}

func (u *userWorkSpaceHandler) Create(c echo.Context) error {
	id := c.Param("workspaceId")

	workSpaceId, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid Workspace ID!",
		})
	}

	var req struct {
		UserName string `json:"userName"`
		Role     string `json:"role"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	repository := repositories.NewUserRepository()

	user, err := repository.GetUserByUserName(req.UserName)

	if err != nil {
		return handleDBError(err, c)
	}

	userWorkSpace := models.UserWorkSpace{
		WorkSpaceID: uint(workSpaceId),
		Role:        req.Role,
		UserID:      user.ID,
	}

	_, err = u.repository.CreateUserWorkSpace(userWorkSpace)
	if err != nil {
		return handleDBError(err, c)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": req.UserName + " Successfully added",
	})
}

func (u *userWorkSpaceHandler) Delete(c echo.Context) error {
	workSpaceID, err := strconv.ParseUint(c.Param("workspaceId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid workspace ID",
		})
	}

	userID, err := strconv.ParseUint(c.Param("userId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid user ID",
		})
	}

	err = u.repository.DeleteUserWorkSpace(uint(workSpaceID), uint(userID))
	if err != nil {
		return handleDBError(err, c)
	}

	return c.NoContent(http.StatusNoContent)
}

func (u *userWorkSpaceHandler) Get(c echo.Context) error {
	panic("not implmented!")
}

func (u *userWorkSpaceHandler) GetAll(c echo.Context) error {
	workSpaceID, err := strconv.ParseUint(c.Param("workspaceId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid workspace ID",
		})
	}

	users, err := u.repository.GetUsersWorkSpace(uint(workSpaceID))
	if err != nil {
		return handleDBError(err, c)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"users": users,
	})
}

func (u *userWorkSpaceHandler) Update(c echo.Context) error {
	workSpaceID, err := strconv.ParseUint(c.Param("workspaceId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid workspace ID",
		})
	}

	userID, err := strconv.ParseUint(c.Param("userId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid user ID",
		})
	}

	var userWorkSpace models.UserWorkSpace
	if err := c.Bind(&userWorkSpace); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request payload",
		})
	}

	userWorkSpace.UserID = uint(userID)
	userWorkSpace.WorkSpaceID = uint(workSpaceID)

	if err := u.repository.UpdateUserWorkSpace(userWorkSpace); err != nil {
		return handleDBError(err, c)
	}

	return c.NoContent(http.StatusNoContent)
}
