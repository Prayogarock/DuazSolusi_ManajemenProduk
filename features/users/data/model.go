package data

import (
	_item "duaz/features/items/data"
	"duaz/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"name;not null"`
	Email    string `gorm:"email;not null"`
	Password string `gorm:"password;not null"`
	Alamat   string `gorm:"alamat;not null"`
	Items    []_item.Item
}

func UserCoreToModel(input users.UserCore) User {
	var userModel = User{
		Model:    gorm.Model{},
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Alamat:   input.Alamat,
	}
	return userModel
}

func UserModelToCore(input User) users.UserCore {
	var userCore = users.UserCore{
		ID:       input.ID,
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Alamat:   input.Alamat,
	}
	return userCore
}
