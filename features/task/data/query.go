package data

import (
	"my-task-api/features/task"

	"gorm.io/gorm"
)

type taskQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) task.DataInterface {
	return &taskQuery{
		db: db,
	}
}

func (p *taskQuery) Insert(input task.Core) error {
	projectGorm := Task{
		UserID:    input.UserID,
		ProjectID: input.ProjectID,
		TaskName:  input.TaskName,
		Status:    input.Status,
	}
	tx := p.db.Create(&projectGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (p *taskQuery) SelectAll(userid uint, projectid uint) ([]task.Core, error) {
	var allProject []Task // var penampung data yg dibaca dari db
	tx := p.db.Where("user_id = ? AND project_id = ?", userid, projectid).Find(&allProject)
	if tx.Error != nil {
		return nil, tx.Error
	}
	//mapping
	var allProjectCore []task.Core
	for _, v := range allProject {
		allProjectCore = append(allProjectCore, task.Core{
			ID:        v.ID,
			UserID:    v.UserID,
			ProjectID: v.ProjectID,
			TaskName:  v.TaskName,
			Status:    v.Status,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	return allProjectCore, nil
}

func (p *taskQuery) Delete(id uint, userid uint) error {
	tx := p.db.Where("id = ? AND user_id = ?", id, userid).Delete(&Task{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (p *taskQuery) PutById(id uint, userid uint, input task.Core) error {

	inputGorm := Task{
		Status: input.Status,
	}
	tx := p.db.Model(&Task{}).Where("id = ? AND user_id = ?", id, userid).Updates(&inputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
