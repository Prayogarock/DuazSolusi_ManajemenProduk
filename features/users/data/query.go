package data

import (
	"duaz/features/users"
	"duaz/helpers"
	"errors"

	"gorm.io/gorm"
)

type UserQuery struct {
	db        *gorm.DB
	dataLogin users.UserCore
}

func New(db *gorm.DB) users.UserDataInterface {
	return &UserQuery{
		db: db,
	}
}

func (repo *UserQuery) Insert(input users.UserCore) error {

	var userModel = UserCoreToModel(input)
	hash, errHass := helpers.HassPassword(userModel.Password)
	if errHass != nil {
		return errHass
	}
	userModel.Password = hash

	tx := repo.db.Create(&userModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}

	return nil
}

func (repo *UserQuery) Login(email string, password string) (dataLogin users.UserCore, err error) {

	var data User
	tx := repo.db.Where("email = ?", email).Find(&data)
	if tx.Error != nil {
		return users.UserCore{}, tx.Error
	}
	check := helpers.CheckPassword(password, data.Password)
	if !check {
		return users.UserCore{}, errors.New("password incorect")
	}
	if tx.RowsAffected == 0 {
		return users.UserCore{}, errors.New("no row affected")
	}
	dataLogin = UserModelToCore(data)
	repo.dataLogin = dataLogin
	return dataLogin, nil
}

func (repo *UserQuery) Read(UserID uint) ([]users.UserCore, error) {
	var userData []User

	tx := repo.db.Where("id = ?", UserID)
	tx.Find(&userData)

	var userCore []users.UserCore
	for _, value := range userData {
		userCore = append(userCore, UserModelToCore(value))
	}

	return userCore, nil
}

func (repo *UserQuery) Update(UserID uint, input users.UserCore) error {
	var user User
	tx := repo.db.Where("id = ?", UserID).First(&user)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("target not found")
	}

	updatedUser := UserCoreToModel(input)

	hashedPassword, err := helpers.HassPassword(updatedUser.Password)
	if err != nil {
		return err
	}
	updatedUser.Password = hashedPassword

	tx = repo.db.Model(&user).Updates(updatedUser)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + " failed to update data")
	}
	return nil
}
