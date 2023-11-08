package users

import (
	_itemCore "duaz/features/items"
)

type UserCore struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Alamat   string
	Items    []_itemCore.ItemCore
}

type UserDataInterface interface {
	Insert(input UserCore) error
	Login(email, password string) (UserCore, error)
	Read(UserID uint) ([]UserCore, error)
	Update(UserID uint, input UserCore) error
}

type UserServiceInterface interface {
	Add(input UserCore) error
	Login(email, password string) (UserCore, string, error)
	GetData(UserID uint) ([]UserCore, error)
	UpdateData(UserID uint, input UserCore) error
}
