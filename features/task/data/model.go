package data

import (
	_dataUser "my-task-api/features/user/data"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	//UserID    uint
	ProjectID uint
	TaskName  string
	Status    string
	User      _dataUser.User
}
