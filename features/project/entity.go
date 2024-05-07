package project

import (
	"my-task-api/features/task"
	"time"
)

type Core struct {
	ID          uint
	UserID      uint
	ProjectName string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Task        []task.Core
}

type DataInterface interface {
	Insert(input Core) error
	SelectAll(userid uint) ([]Core, error)
	// SelectByUsername(username string) (*Core, error)
	Delete(id uint, userid uint) error
	// PutToken(username string, input string) error
	//SelectById(id uint) (*Core, error)
	PutById(id uint, userid uint, input Core) error
}

type ServiceInterface interface {
	Create(input Core) error
	GetAll(userid uint) ([]Core, error)
	// Login(username, password string) (data *Core, token string, err error)
	Delete(id uint, userid uint) error
	// UpdateToken(username string, input string) error
	//GetById(id uint) (data *Core, err error)
	UpdateById(id uint, userid uint, input Core) error
}
