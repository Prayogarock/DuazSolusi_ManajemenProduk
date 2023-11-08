package handler

import "duaz/features/users"

type LoginResponse struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type UserResponseAll struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Alamat string `json:"alamat"`
}

func UserCoreToResponseAll(input users.UserCore) UserResponseAll {
	var resultResponse = UserResponseAll{
		ID:     input.ID,
		Name:   input.Name,
		Email:  input.Email,
		Alamat: input.Alamat,
	}
	return resultResponse
}
