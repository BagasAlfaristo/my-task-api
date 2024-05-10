package handler

import (
	"my-task-api/features/task"
	"my-task-api/utils/responses"
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
	newProject := TaskAddRequest{}
	errBind := c.Bind(&newProject)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind data: "+errBind.Error(), nil))
	}

	//idToken := middlewares.ExtractTokenUserId(c)
	//hashedPassword := hashPassword(newUser.Password)
	//newUser.Password = hashedPassword
	// mapping  dari request ke core
	Default := "Not Completed"
	inputCore := task.Core{
		ProjectID: newProject.ProjectID,
		TaskName:  newProject.TaskName,
		Status:    Default,
	}

	// memanggil/mengirimkan data ke method service layer
	errInsert := ph.taskService.Create(inputCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error add data", errInsert))
	}
	return c.JSON(http.StatusCreated, responses.JSONWebResponse("success add data", errInsert))
}

func (uh *TaskHandler) GetAll(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error get id", idConv))
	}

	result, err := uh.taskService.GetAll(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error read data", result))
	}
	var allUsersResponse []TaskResponse
	for _, value := range result {
		allUsersResponse = append(allUsersResponse, TaskResponse{
			ID:       value.ID,
			TaskName: value.TaskName,
			Status:   value.Status,
		})
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success read data", allUsersResponse))
}

func (uh *TaskHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error get id ", idConv))
	}

	err := uh.taskService.Delete(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error delete data", err))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success delete data", err))
}

func (uh *TaskHandler) UpdateById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error get id", idConv))
	}

	updatedProject := TaskRequest{}
	errBind := c.Bind(&updatedProject)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind"+errBind.Error(), nil))
	}

	// mapping  dari request ke core
	inputNewCore := task.Core{
		Status: updatedProject.Status,
	}

	err := uh.taskService.UpdateById(uint(idConv), inputNewCore)
	if err != nil {
		// Handle error from userService.UpdateById
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error update data", err))
	}

	// Return success response
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success update data", err))
}
