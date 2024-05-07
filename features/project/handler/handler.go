package handler

import (
	"log"
	"my-task-api/app/middlewares"
	"my-task-api/features/project"
	"net/http"
	"strconv"

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
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error bind data: " + errBind.Error(),
		})
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

func (uh *ProjectHandler) GetAll(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	log.Println("idtoken:", idToken)
	result, err := uh.projectService.GetAll(uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error read data",
		})
	}
	var allUsersResponse []ProjectResponse
	for _, value := range result {
		allUsersResponse = append(allUsersResponse, ProjectResponse{
			ID:          value.ID,
			ProjectName: value.ProjectName,
			Description: value.Description,
		})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success read data",
		"results": allUsersResponse,
	})
}

func (uh *ProjectHandler) UpdateById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error convert id: " + errConv.Error(),
		})
	}

	idToken := middlewares.ExtractTokenUserId(c)
	updatedProject := ProjectRequest{}
	errBind := c.Bind(&updatedProject)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "error bind data: " + errBind.Error(),
		})
	}

	// mapping  dari request ke core
	inputNewCore := project.Core{
		ProjectName: updatedProject.ProjectName,
		Description: updatedProject.Description,
	}

	err := uh.projectService.UpdateById(uint(idConv), uint(idToken), inputNewCore)
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

func (uh *ProjectHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error convert id: " + errConv.Error(),
		})
	}

	idToken := middlewares.ExtractTokenUserId(c)
	err := uh.projectService.Delete(uint(idToken), uint(idConv))
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
