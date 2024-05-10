package handler

import (
	"log"
	"my-task-api/app/middlewares"
	"my-task-api/features/user"
	"my-task-api/utils/responses"

	//user "my-task-api/features"
	"net/http"

	//"github.com/labstack/echo"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.ServiceInterface
}

func New(us user.ServiceInterface) *UserHandler {
	return &UserHandler{
		userService: us,
	}
}

func (uh *UserHandler) Register(c echo.Context) error {
	// membaca data dari request body
	newUser := UserRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind"+errBind.Error(), nil))
	}

	//hashedPassword := hashPassword(newUser.Password)
	//newUser.Password = hashedPassword
	// mapping  dari request ke core
	inputCore := user.Core{
		Name:        newUser.Name,
		Email:       newUser.Email,
		Username:    newUser.Username,
		Password:    newUser.Password,
		PhoneNumber: newUser.PhoneNumber,
	}

	// memanggil/mengirimkan data ke method service layer
	errInsert := uh.userService.Create(inputCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error add data", errInsert))
	}
	return c.JSON(http.StatusCreated, responses.JSONWebResponse("success add data", errInsert))
}

func (uh *UserHandler) GetAll(c echo.Context) error {
	result, err := uh.userService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error read data", result))
	}
	var allUsersResponse []UserResponse
	for _, value := range result {
		allUsersResponse = append(allUsersResponse, UserResponse{
			ID:          value.ID,
			Name:        value.Name,
			Email:       value.Email,
			Username:    value.Username,
			PhoneNumber: value.PhoneNumber,
		})
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success read data", allUsersResponse))
}

func (uh *UserHandler) Login(c echo.Context) error {
	var reqLoginData = LoginRequest{}
	errBind := c.Bind(&reqLoginData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind"+errBind.Error(), nil))
	}
	result, token, err := uh.userService.Login(reqLoginData.Username, reqLoginData.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error login", result))
	}
	//mapping
	var resultResponse = map[string]any{
		"id":    result.ID,
		"name":  result.Name,
		"token": token,
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success login", resultResponse))
}

func (uh *UserHandler) Profile(c echo.Context) error {
	// extract id user from jwt token
	idToken := middlewares.ExtractTokenUserId(c)
	log.Println("idtoken:", idToken)
	result, err := uh.userService.GetById(uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error get id", result))
	}
	resultResponse := UserResponse{
		ID:    result.ID,
		Name:  result.Name,
		Email: result.Email,
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success get profile", resultResponse))
}

func (uh *UserHandler) Delete(c echo.Context) error {
	//id := c.Param("id")
	idToken := middlewares.ExtractTokenUserId(c)
	//idConv, errConv := strconv.Atoi(idToken)
	// if errConv != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]any{
	// 		"status":  "failed",
	// 		"message": "error convert id: " + errConv.Error(),
	// 	})
	// }
	err := uh.userService.Delete(uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error delete data", err))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success delete data", err))
}

func (uh *UserHandler) UpdateById(c echo.Context) error {
	// id := c.Param("id")
	// idConv, errConv := strconv.Atoi(id)
	idToken := middlewares.ExtractTokenUserId(c)

	updatedUser := UserRequest{}
	errBind := c.Bind(&updatedUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind"+errBind.Error(), nil))
	}

	// mapping  dari request ke core
	inputNewCore := user.Core{
		Name:        updatedUser.Name,
		Email:       updatedUser.Email,
		Username:    updatedUser.Username,
		Password:    updatedUser.Password,
		PhoneNumber: updatedUser.PhoneNumber,
	}

	err := uh.userService.UpdateById(uint(idToken), inputNewCore)
	if err != nil {
		// Handle error from userService.UpdateById
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error update data", err))
	}

	// Return success response
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success update data", err))
}
