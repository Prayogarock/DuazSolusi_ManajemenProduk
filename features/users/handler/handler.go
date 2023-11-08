package handler

import (
	"duaz/app/middlewares"
	"duaz/features/users"
	"duaz/helpers"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService users.UserServiceInterface
}

func New(service users.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) Add(c echo.Context) error {
	var userInput UserRequest
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid"+errBind.Error(), nil))
	}

	validate := validator.New()
	if err := validate.Struct(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
	}
	var userCore = UserRequestToCore(userInput)
	err := handler.userService.Add(userCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusCreated, "success register user", nil))
}

func (handler *UserHandler) Login(c echo.Context) error {
	userInput := new(LoginRequest)
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	dataLogin, token, err := handler.userService.Login(userInput.Email, userInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error login", nil))

		}
	}
	var response = LoginResponse{
		Name:  dataLogin.Name,
		Token: token,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success login", response))
}

func (handler *UserHandler) GetUserData(c echo.Context) error {
	UserID, err := middlewares.ExtractTokenUser(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, err.Error(), nil))
	}

	result, err := handler.userService.GetData(UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}

	var userResponse []UserResponseAll
	for _, value := range result {
		userResponse = append(userResponse, UserResponseAll{
			ID:     value.ID,
			Name:   value.Name,
			Email:  value.Email,
			Alamat: value.Alamat,
		})
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", userResponse))
}

func (handler *UserHandler) UpdateUser(c echo.Context) error {
	UserID, err := middlewares.ExtractTokenUser(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, err.Error(), nil))
	}

	var userInput UserRequest
	if errBind := c.Bind(&userInput); errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. Data not valid", nil))
	}

	updatedUser := UserRequestToCore(userInput)
	err = handler.userService.UpdateData(UserID, updatedUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error updating user data", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success update user data", nil))
}
