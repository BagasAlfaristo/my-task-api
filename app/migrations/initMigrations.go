package migrations

import (
	"gorm.io/gorm"
)

type migrationsQuery struct {
	db *gorm.DB
}

type User struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Email       string `gorm:"unique" json:"email" form:"email"`
	Username    string `gorm:"unique" json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
	PhoneNumber string `gorm:"unique" json:"phonenumber" form:"phonenumber"`
	//Products  []Product `gorm:"foreignKey:UserID;references:ID"`
	// Favorites []Favorite
}

type Project struct {
	gorm.Model
	UserID      uint
	ProjectName string `json:"name" form:"name"`
	Description string
	User        User   `gorm:"foreignKey:UserID;references:ID"`
	Task        []Task `gorm:"foreignKey:ProjectID;references:ID"`
	// Favorites []Favorite
}

type Task struct {
	gorm.Model
	UserID    uint
	ProjectID uint
	TaskName  string
	Status    string
	User      User    `gorm:"foreignKey:UserID;references:ID"`
	Project   Project `gorm:"foreignKey:ProjectID;references:ID"`
}

func InitMigrations(m *migrationsQuery) {
	m.db.AutoMigrate(&User{})
	m.db.AutoMigrate(&Project{})
	m.db.AutoMigrate(&Task{})
}
