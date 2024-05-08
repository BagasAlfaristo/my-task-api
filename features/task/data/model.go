package data

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ProjectID uint
	TaskName  string
	Status    string
}
