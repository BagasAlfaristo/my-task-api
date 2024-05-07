package handler

import "my-task-api/features/task"

type ProjectResponse struct {
	ID          uint        `json:"id"`
	ProjectName string      `json:"projectname"`
	Description string      `json:"description"`
	Task        []task.Core `gorm:"foreignKey:ProjectID;references:ID"`
}
