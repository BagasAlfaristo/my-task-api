package handler

import "gorm.io/gorm"

type ProjectRequest struct {
	gorm.Model
	UserID      uint   `json:"userid" form:"userid"`
	ProjectName string `json:"projectname" form:"projectname"`
	Description string `json:"description" form:"description"`
}
