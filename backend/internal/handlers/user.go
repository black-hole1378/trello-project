package handlers

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	repository repositories.UserRepoImpl
}

func NewUserHandler() HandlerImpl {
	return &userHandler{
		repository: repositories.NewUserRepository(),
	}
}

func (t *userHandler) Update(c echo.Context) error {
	userName, userID, err := getUsername(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if user.Email != "" {
		if err := utils.IsValidEmail(user.Email); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}
	}

	if user.Password != "" {
		if err := utils.PasswordCheck(user.Password); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}
	}

	if err = t.repository.UpdateUser(&user, userName, userID); err != nil {
		return handleDBError(err, c)
	}

	accessToken, refreshToken, err := utils.GenerateJWT(userName, c.Get("password").(string))

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message":      "Successfully Updated!",
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})

}

func (t *userHandler) GetAll(c echo.Context) error {
	users, err := t.repository.GetAllUsers()

	if err != nil {
		return handleDBError(err, c)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"users": users,
	})

}

func getUsername(c echo.Context) (string, uint, error) {
	id := c.Param("userID")

	userName := c.Get("userName").(string)

	userID, err := strconv.ParseUint(id, 10, 32)

	return userName, uint(userID), err
}

func (t *userHandler) Get(c echo.Context) error {

	userName, userID, err := getUsername(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	user, err := t.repository.GetUser(userID, userName)

	if err != nil {
		return handleDBError(err, c)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user": user,
	})
}

func (t *userHandler) Create(c echo.Context) error {
	var newUser models.User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := utils.IsValidEmail(newUser.Email); err != nil {

		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := utils.PasswordCheck(newUser.Password); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	newUser.Password = utils.HashedPassword(newUser.Password)

	user, err := t.repository.CreateUser(&newUser)
     
	if err != nil {
		return handleDBError(err, c)
	}
    
	accessToken, refreshToken, _ := utils.GenerateJWT(user.Username, user.Password)
	
	return c.JSON(http.StatusCreated, echo.Map{
		"user":         user,
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func (t *userHandler) Delete(c echo.Context) error {
	userName, userID, err := getUsername(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	if err := t.repository.DeleteUser(userID, userName); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Something Bad happened!",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Deleted Successfully!",
	})
}
