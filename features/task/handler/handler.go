package handler

import (
	"log"
	"my-task-api/app/middlewares"
	"my-task-api/features/task"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskService task.ServiceInterface
}

func New(ts task.ServiceInterface) *TaskHandler {
	return &TaskHandler{
		taskService: ts,
	}
}

func (ph *TaskHandler) Register(c echo.Context) error {
	// membaca data dari request body
	newProject := TaskRequest{}
	errBind := c.Bind(&newProject)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error bind data: " + errBind.Error(),
		})
	}

	idToken := middlewares.ExtractTokenUserId(c)
	//hashedPassword := hashPassword(newUser.Password)
	//newUser.Password = hashedPassword
	// mapping  dari request ke core
	inputCore := task.Core{
		ID:        newProject.ID,
		UserID:    uint(idToken),
		ProjectID: newProject.ProjectID,
		TaskName:  newProject.TaskName,
		Status:    newProject.Status,
	}

	// memanggil/mengirimkan data ke method service layer
	errInsert := ph.taskService.Create(inputCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error insert data " + errInsert.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]any{
		"status":  "success",
		"message": "success add data",
	})
}

func (uh *TaskHandler) GetAll(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error convert id: " + errConv.Error(),
		})
	}

	idToken := middlewares.ExtractTokenUserId(c)
	log.Println("idtoken:", idToken)
	result, err := uh.taskService.GetAll(uint(idToken), uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error read data",
		})
	}
	var allUsersResponse []TaskResponse
	for _, value := range result {
		allUsersResponse = append(allUsersResponse, TaskResponse{
			ID:       value.ID,
			TaskName: value.TaskName,
			Status:   value.Status,
		})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success read data",
		"results": allUsersResponse,
	})
}

func (uh *TaskHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error convert id: " + errConv.Error(),
		})
	}

	idToken := middlewares.ExtractTokenUserId(c)
	err := uh.taskService.Delete(uint(idConv), uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error delete data " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success delete data",
	})
}

func (uh *TaskHandler) UpdateById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error convert id: " + errConv.Error(),
		})
	}

	idToken := middlewares.ExtractTokenUserId(c)
	updatedProject := TaskRequest{}
	errBind := c.Bind(&updatedProject)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "error bind data: " + errBind.Error(),
		})
	}

	// mapping  dari request ke core
	inputNewCore := task.Core{
		Status: updatedProject.Status,
	}

	err := uh.taskService.UpdateById(uint(idConv), uint(idToken), inputNewCore)
	if err != nil {
		// Handle error from userService.UpdateById
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": "error updating project by id: " + err.Error(),
		})
	}

	// Return success response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "data updated successfully",
	})
}
