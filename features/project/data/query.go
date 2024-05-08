package data

import (
	"my-task-api/features/project"
	"my-task-api/features/task"

	"gorm.io/gorm"
)

type projectQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) project.DataInterface {
	return &projectQuery{
		db: db,
	}
}

func (p *projectQuery) Insert(input project.Core) error {
	projectGorm := Project{
		UserID:      input.UserID,
		ProjectName: input.ProjectName,
		Description: input.Description,
	}
	tx := p.db.Create(&projectGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (p *projectQuery) SelectAll(userid uint) ([]project.Core, error) {
	var allProject []Project // var penampung data yg dibaca dari db
	tx := p.db.Where("user_id = ?", userid).Preload("Tasks").Find(&allProject)
	if tx.Error != nil {
		return nil, tx.Error
	}
	//mapping
	var allProjectCore []project.Core
	for _, v := range allProject {
		var allTaskCore []task.Core
		for _, vtask := range v.Tasks {
			allTaskCore = append(allTaskCore, task.Core{
				ID:        vtask.ID,
				ProjectID: vtask.ProjectID,
				TaskName:  vtask.TaskName,
				Status:    vtask.Status,
				CreatedAt: vtask.CreatedAt,
				UpdatedAt: vtask.UpdatedAt,
			})
		}
		allProjectCore = append(allProjectCore, project.Core{
			ID:          v.ID,
			ProjectName: v.ProjectName,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			Task:        allTaskCore,
		})
	}
	return allProjectCore, nil
}

func (p *projectQuery) PutById(id uint, userid uint, input project.Core) error {

	inputGorm := Project{
		ProjectName: input.ProjectName,
		Description: input.Description,
	}
	tx := p.db.Model(&Project{}).Where("id = ? AND user_id = ?", id, userid).Updates(&inputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (p *projectQuery) Delete(id uint, userid uint) error {
	tx := p.db.Where("user_id = ?", userid).Delete(&Project{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
