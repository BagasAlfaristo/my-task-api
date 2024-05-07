package data

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Email       string `gorm:"unique" json:"email" form:"email"`
	Username    string `gorm:"unique" json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
	PhoneNumber string `gorm:"unique" json:"phonenumber" form:"phonenumber"`
	Token       string `json:"token" form:"token"`
	//Products  []Product `gorm:"foreignKey:UserID;references:ID"`
	// Favorites []Favorite
}
