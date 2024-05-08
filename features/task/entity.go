package task

import "time"

type Core struct {
	ID        uint
	ProjectID uint
	TaskName  string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DataInterface interface {
	Insert(input Core) error
	SelectAll(projecid uint) ([]Core, error)
	// // SelectByUsername(username string) (*Core, error)
	Delete(id uint) error
	// // PutToken(username string, input string) error
	// //SelectById(id uint) (*Core, error)
	PutById(id uint, input Core) error
}

type ServiceInterface interface {
	Create(input Core) error
	GetAll(projecid uint) ([]Core, error)
	// // Login(username, password string) (data *Core, token string, err error)
	Delete(id uint) error
	// // UpdateToken(username string, input string) error
	// //GetById(id uint) (data *Core, err error)
	UpdateById(id uint, input Core) error
}
