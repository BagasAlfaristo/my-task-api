package service

import (
	"errors"
	"my-task-api/features/project"
)

type projectService struct {
	projectData project.DataInterface
	//projectService encrypts.HashInterface
}

func New(pd project.DataInterface) project.ServiceInterface {
	return &projectService{
		projectData: pd,
		//hashService: hash,
	}

}

func (p *projectService) Create(input project.Core) error {
	err := p.projectData.Insert(input)
	if err != nil {
		return err
	}
	return nil
}

func (p *projectService) GetAll(userid uint) ([]project.Core, error) {
	if userid <= 0 {
		return nil, errors.New("[validation] id not valid")
	}
	return p.projectData.SelectAll(userid)
}

func (p *projectService) UpdateById(id uint, userid uint, input project.Core) error {
	if id <= 0 {
		return errors.New("id not valid")
	}
	err := p.projectData.PutById(id, userid, input)
	if err != nil {
		return err
	}
	return nil
}

func (u *projectService) Delete(id uint, userid uint) error {
	if id <= 0 {
		return errors.New("id not valid")
	}
	return u.projectData.Delete(id, userid)
}
