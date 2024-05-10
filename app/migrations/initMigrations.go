package migrations

import (
	_dataProject "my-task-api/features/project/data"
	_dataTask "my-task-api/features/task/data"
	_dataUser "my-task-api/features/user/data"

	"gorm.io/gorm"
)

func InitMigrations(db *gorm.DB) {
	db.AutoMigrate(&_dataUser.User{})
	db.AutoMigrate(&_dataProject.Project{})
	db.AutoMigrate(&_dataTask.Task{})
}
