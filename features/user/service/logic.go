package service

import (
	"errors"
	"my-task-api/app/middlewares"
	"my-task-api/features/user"
	"my-task-api/utils/encrypts"
)

type userService struct {
	userData    user.DataInterface
	hashService encrypts.HashInterface
}

func New(ud user.DataInterface, hash encrypts.HashInterface) user.ServiceInterface {
	return &userService{
		userData:    ud,
		hashService: hash,
	}

}

func (u *userService) Create(input user.Core) error {
	if input.Name == "" || input.Email == "" || input.Password == "" {
		return errors.New("nama/email/password tidak boleh kosong")
	}

	if input.Password != "" {
		// proses hash password
		result, errHash := u.hashService.HashPassword(input.Password)
		if errHash != nil {
			return errHash
		}
		input.Password = result
	}

	err := u.userData.Insert(input)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) GetAll() ([]user.Core, error) {
	return u.userData.SelectAll()
}

// Login implements user.ServiceInterface.
func (u *userService) Login(username string, password string) (data *user.Core, token string, err error) {
	data, err = u.userData.SelectByUsername(username)
	if err != nil {
		return nil, "", err
	}

	isLoginValid := u.hashService.CheckPasswordHash(data.Password, password)
	// ketika isloginvalid = true, maka login berhasil
	if !isLoginValid {
		return nil, "", errors.New("[validation] password tidak sesuai")
	}
	token, errJWT := middlewares.CreateToken(int(data.ID))
	if errJWT != nil {
		return nil, "", errJWT
	}

	u.UpdateToken(username, token)
	return data, token, nil
}

func (u *userService) UpdateToken(username string, token string) error {

	err := u.userData.PutToken(username, token)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) GetById(id uint) (data *user.Core, err error) {
	if id <= 0 {
		return nil, errors.New("[validation] id not valid")
	}
	return u.userData.SelectById(id)
}

func (u *userService) Delete(id uint) error {
	if id <= 0 {
		return errors.New("id not valid")
	}
	return u.userData.Delete(id)
}

func (u *userService) UpdateById(id uint, input user.Core) error {
	if id <= 0 {
		return errors.New("id not valid")
	} else if input.Name == "" || input.Email == "" || input.Password == "" {
		return errors.New("nama/email/password tidak boleh kosong")
	}

	err := u.userData.PutById(id, input)
	if err != nil {
		return err
	}
	return nil
}
