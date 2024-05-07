package handler

import (
	"log"
	"my-task-api/app/middlewares"
	"my-task-api/features/user"

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
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error bind data: " + errBind.Error(),
		})
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
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error insert data " + errInsert.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]any{
		"status":  "success",
		"message": "success add user",
	})
}

func (uh *UserHandler) GetAll(c echo.Context) error {
	result, err := uh.userService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error read data",
		})
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
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success read data",
		"results": allUsersResponse,
	})
}

func (uh *UserHandler) Login(c echo.Context) error {
	var reqLoginData = LoginRequest{}
	errBind := c.Bind(&reqLoginData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error bind data: " + errBind.Error(),
		})
	}
	result, token, err := uh.userService.Login(reqLoginData.Username, reqLoginData.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error login " + err.Error(),
		})
	}
	//mapping
	var resultResponse = map[string]any{
		"id":    result.ID,
		"name":  result.Name,
		"token": token,
	}
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success login",
		"data":    resultResponse,
	})
}

func (uh *UserHandler) Profile(c echo.Context) error {
	// extract id user from jwt token
	idToken := middlewares.ExtractTokenUserId(c)
	log.Println("idtoken:", idToken)
	result, err := uh.userService.GetById(uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error login " + err.Error(),
		})
	}
	resultResponse := UserResponse{
		ID:    result.ID,
		Name:  result.Name,
		Email: result.Email,
	}
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success login",
		"data":    resultResponse,
	})
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
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error delete data " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success delete user",
	})
}

func (uh *UserHandler) UpdateById(c echo.Context) error {
	// id := c.Param("id")
	// idConv, errConv := strconv.Atoi(id)
	idToken := middlewares.ExtractTokenUserId(c)

	updatedUser := UserRequest{}
	errBind := c.Bind(&updatedUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "error bind data: " + errBind.Error(),
		})
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
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": "error updating user by id: " + err.Error(),
		})
	}

	// Return success response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "user updated successfully",
	})
}
