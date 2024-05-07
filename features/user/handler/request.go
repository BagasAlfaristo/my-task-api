package handler

import "gorm.io/gorm"

type UserRequest struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Email       string `gorm:"unique" json:"email" form:"email"`
	Username    string `gorm:"unique" json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
	PhoneNumber string `gorm:"unique" json:"phonenumber" form:"phonenumber"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
