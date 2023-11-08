package handler

import "duaz/features/users"

type UserRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Alamat   string `json:"alamat" form:"alamat" validate:"required"`
}

func UserRequestToCore(input UserRequest) users.UserCore {
	var userCore = users.UserCore{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Alamat:   input.Alamat,
	}
	return userCore
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
