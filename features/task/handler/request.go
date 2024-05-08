package handler

type TaskRequest struct {
	ProjectID uint   `json:"project_id" form:"project_id"`
	TaskName  string `json:"task_name" form:"task_name"`
	Status    string `json:"status" form:"status"`
}

type TaskAddRequest struct {
	ProjectID uint   `json:"project_id" form:"project_id"`
	TaskName  string `json:"task_name" form:"task_name"`
	Status    string `json:"status" form:"status"`
}
