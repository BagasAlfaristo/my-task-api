package handler

import "gorm.io/gorm"

type TaskRequest struct {
	gorm.Model
	ProjectID uint   `json:"project_id" form:"project_id"`
	UserID    uint   `json:"user_id" form:"user_id"`
	TaskName  string `json:"task_name" form:"task_name"`
	Status    string `json:"status" form:"status"`
}
