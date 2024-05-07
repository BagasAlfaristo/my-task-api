package service

import (
	"errors"
	"my-task-api/features/task"
)

type taskService struct {
	taskData task.DataInterface
	//projectService encrypts.HashInterface
}

func New(pd task.DataInterface) task.ServiceInterface {
	return &taskService{
		taskData: pd,
		//hashService: hash,
	}

}

func (p *taskService) Create(input task.Core) error {
	err := p.taskData.Insert(input)
	if err != nil {
		return err
	}
	return nil
}

func (p *taskService) GetAll(userid uint, projectid uint) ([]task.Core, error) {
	if userid <= 0 {
		return nil, errors.New("[validation] id not valid")
	}
	return p.taskData.SelectAll(userid, projectid)
}

func (u *taskService) Delete(id uint, userid uint) error {
	if id <= 0 {
		return errors.New("id not valid")
	}
	return u.taskData.Delete(id, userid)
}

func (p *taskService) UpdateById(id uint, userid uint, input task.Core) error {
	if id <= 0 {
		return errors.New("id not valid")
	}
	err := p.taskData.PutById(id, userid, input)
	if err != nil {
		return err
	}
	return nil
}
