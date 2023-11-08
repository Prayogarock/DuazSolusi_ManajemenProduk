package service

import (
	"duaz/app/middlewares"
	"duaz/features/users"
)

type UserService struct {
	userData users.UserDataInterface
}

func New(repo users.UserDataInterface) users.UserServiceInterface {
	return &UserService{
		userData: repo,
	}
}

func (service *UserService) Add(input users.UserCore) error {
	return service.userData.Insert(input)
}

func (service *UserService) Login(email string, password string) (dataLogin users.UserCore, token string, err error) {

	dataLogin, err = service.userData.Login(email, password)
	if err != nil {
		return users.UserCore{}, "", err
	}
	token, err = middlewares.CreateTokenUser(dataLogin.ID)
	if err != nil {
		return users.UserCore{}, "", err
	}
	return dataLogin, token, nil
}

func (service *UserService) GetData(UserID uint) ([]users.UserCore, error) {
	return service.userData.Read(UserID)
}

func (service *UserService) UpdateData(UserID uint, input users.UserCore) error {
	return service.userData.Update(UserID, input)
}
