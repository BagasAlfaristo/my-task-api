package data

import (
	"my-task-api/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.DataInterface {
	return &userQuery{
		db: db,
	}
}

// Insert implements user.DataInterface.
func (u *userQuery) Insert(input user.Core) error {

	userGorm := User{
		Name:        input.Name,
		Email:       input.Email,
		Username:    input.Username,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
	}
	tx := u.db.Create(&userGorm)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// SelectAll implements user.DataInterface.
func (u *userQuery) SelectAll() ([]user.Core, error) {
	var allUsers []User // var penampung data yg dibaca dari db
	tx := u.db.Find(&allUsers)
	if tx.Error != nil {
		return nil, tx.Error
	}
	//mapping
	var allUserCore []user.Core
	for _, v := range allUsers {
		allUserCore = append(allUserCore, user.Core{
			ID:          v.ID,
			Name:        v.Name,
			Email:       v.Email,
			Username:    v.Username,
			Password:    v.Password,
			PhoneNumber: v.PhoneNumber,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}

	return allUserCore, nil
}

// Login implements user.DataInterface.
func (u *userQuery) SelectByUsername(username string) (*user.Core, error) {
	// select id, name, phone from users where email = xxxx and password = xxxxxx
	// select id, name, phone from users where email = xxxx

	// variable penampung datanya
	var userData User
	tx := u.db.Where("username = ?", username).First(&userData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// mapping
	var usercore = user.Core{
		ID:          userData.ID,
		Name:        userData.Name,
		Username:    userData.Username,
		Email:       userData.Email,
		Password:    userData.Password,
		PhoneNumber: userData.PhoneNumber,
		CreatedAt:   userData.CreatedAt,
		UpdatedAt:   userData.UpdatedAt,
	}

	return &usercore, nil
}

func (u *userQuery) PutToken(username string, token string) error {
	inputGorm := User{
		Token: token,
	}
	tx := u.db.Model(&User{}).Where("username = ?", username).Updates(&inputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (u *userQuery) SelectById(id uint) (*user.Core, error) {
	// variable penampung datanya
	var userData User
	tx := u.db.First(&userData, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// mapping
	var usercore = user.Core{
		ID:          userData.ID,
		Name:        userData.Name,
		Email:       userData.Email,
		Password:    userData.Password,
		PhoneNumber: userData.PhoneNumber,
		Username:    userData.Username,
		CreatedAt:   userData.CreatedAt,
		UpdatedAt:   userData.UpdatedAt,
	}

	return &usercore, nil
}

func (u *userQuery) Delete(id uint) error {
	tx := u.db.Delete(&User{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u *userQuery) PutById(id uint, input user.Core) error {

	inputGorm := User{
		Name:        input.Name,
		Email:       input.Email,
		Username:    input.Username,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
	}
	tx := u.db.Model(&User{}).Where("id = ?", id).Updates(&inputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
