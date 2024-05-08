package data

import (
	"my-task-api/features/task/data"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	UserID      uint
	ProjectName string
	Description string
	Tasks       []data.Task
}
