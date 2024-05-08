package handler

import "my-task-api/features/task/handler"

type ProjectResponse struct {
	ID          uint                   `json:"id"`
	ProjectName string                 `json:"projectname"`
	Description string                 `json:"description"`
	Task        []handler.TaskResponse `gorm:"foreignKey:ProjectID;references:ID"`
}
