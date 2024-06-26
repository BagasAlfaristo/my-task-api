package handler

import (
	"log"
	"my-task-api/app/middlewares"
	"my-task-api/features/project"
	"my-task-api/features/task/handler"
	"my-task-api/utils/responses"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	projectService project.ServiceInterface
}

func New(ps project.ServiceInterface) *ProjectHandler {
	return &ProjectHandler{
		projectService: ps,
	}
}

func (ph *ProjectHandler) Register(c echo.Context) error {
	// membaca data dari request body
	newProject := ProjectRequest{}
	errBind := c.Bind(&newProject)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind data: "+errBind.Error(), nil))
	}

	idToken := middlewares.ExtractTokenUserId(c)
	//hashedPassword := hashPassword(newUser.Password)
	//newUser.Password = hashedPassword
	// mapping  dari request ke core
	inputCore := project.Core{
		ID:          newProject.ID,
		UserID:      uint(idToken),
		ProjectName: newProject.ProjectName,
		Description: newProject.Description,
	}

	// memanggil/mengirimkan data ke method service layer
	errInsert := ph.projectService.Create(inputCore)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error insert data: "+errInsert.Error(), nil))

		}
	}
	return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error insert data: "+errInsert.Error(), nil))
}

func (uh *ProjectHandler) GetAll(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	log.Println("idtoken:", idToken)
	result, err := uh.projectService.GetAll(uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error read data", nil))
	}
	var allUsersResponse []ProjectResponse
	for _, value := range result {
		var allTaskResponse []handler.TaskResponse
		for _, vtask := range value.Task {
			allTaskResponse = append(allTaskResponse, handler.TaskResponse{
				ID:       vtask.ID,
				TaskName: vtask.TaskName,
				Status:   vtask.Status,
			})
		}
		allUsersResponse = append(allUsersResponse, ProjectResponse{
			ID:          value.ID,
			ProjectName: value.ProjectName,
			Description: value.Description,
			Task:        allTaskResponse,
		})
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success read data", allUsersResponse))
}

func (uh *ProjectHandler) UpdateById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error get user id", idConv))
	}

	idToken := middlewares.ExtractTokenUserId(c)
	updatedProject := ProjectRequest{}
	errBind := c.Bind(&updatedProject)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind data: "+errBind.Error(), nil))
	}

	// mapping  dari request ke core
	inputNewCore := project.Core{
		ProjectName: updatedProject.ProjectName,
		Description: updatedProject.Description,
	}

	err := uh.projectService.UpdateById(uint(idConv), uint(idToken), inputNewCore)
	if err != nil {
		// Handle error from userService.UpdateById
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error update data", err))
	}

	// Return success response
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success update data", err))
}

func (uh *ProjectHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error get user id", idConv))
	}

	idToken := middlewares.ExtractTokenUserId(c)
	err := uh.projectService.Delete(uint(idToken), uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error delete data", err))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success delete data", err))
}
