package data

import (
	"my-task-api/features/task"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	UserID      uint
	ProjectName string
	Description string
	Task        []task.Core
}
