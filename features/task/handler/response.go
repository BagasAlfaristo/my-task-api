package handler

type TaskResponse struct {
	ID       uint   `json:"id"`
	TaskName string `json:"task_name" form:"task_name"`
	Status   string `json:"status" form:"status"`
}
