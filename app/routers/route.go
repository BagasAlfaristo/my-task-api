package routers

import (
	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"my-task-api/app/middlewares"
	_userData "my-task-api/features/user/data"
	_userHandler "my-task-api/features/user/handler"
	_userService "my-task-api/features/user/service"
	"my-task-api/utils/encrypts"

	_projectData "my-task-api/features/project/data"
	_projectHandler "my-task-api/features/project/handler"
	_projectService "my-task-api/features/project/service"

	_taskData "my-task-api/features/task/data"
	_taskHandler "my-task-api/features/task/handler"
	_taskService "my-task-api/features/task/service"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {

	hashService := encrypts.NewHashService()
	userData := _userData.New(db)
	userService := _userService.New(userData, hashService)
	userHandlerAPI := _userHandler.New(userService)

	projectData := _projectData.New(db)
	projectService := _projectService.New(projectData)
	projectHandlerAPI := _projectHandler.New(projectService)

	taskData := _taskData.New(db)
	taskService := _taskService.New(taskData)
	taskHandlerAPI := _taskHandler.New(taskService)

	e.GET("/users", userHandlerAPI.GetAll)
	e.POST("/users", userHandlerAPI.Register)
	e.POST("/login", userHandlerAPI.Login)
	e.GET("/profile", userHandlerAPI.Profile, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.Delete)
	e.PUT("/users", userHandlerAPI.UpdateById)

	e.POST("/projects", projectHandlerAPI.Register)
	e.GET("/projects", projectHandlerAPI.GetAll)
	e.PUT("/projects/:id", projectHandlerAPI.UpdateById)
	e.DELETE("/projects/:id", projectHandlerAPI.Delete)

	e.POST("/tasks", taskHandlerAPI.Register)
	e.GET("/tasks/:id", taskHandlerAPI.GetAll)
	e.DELETE("/tasks/:id", taskHandlerAPI.Delete)
	e.PUT("/tasks/:id", taskHandlerAPI.UpdateById)
}
