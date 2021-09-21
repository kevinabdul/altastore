package libdb

import (
	"altastore/config"
	"altastore/models"
	"altastore/util/password"
	"gorm.io/gorm"

	"errors"
	//"fmt"
	"strings"
)

func GetAdminByUserId(targetId int) (models.UserAPI, error) {
	var user models.UserAPI

	res := config.Db.Model(&models.User{}).Find(&user, targetId)

	if res.Error != nil {
		return models.UserAPI{}, res.Error
	}

	if res.RowsAffected == 0 {
		return models.UserAPI{}, errors.New("Wrong User Id")
	}

	return user, nil
}

func AddAdmin(newUser *models.User) (models.UserAPI, error) {
	if newUser.Email == "" || newUser.Password == "" {
		return models.UserAPI{}, errors.New("Invalid Email or Password. Make sure its not empty and are of string type")
	}

	if newUser.Name == "" {
		return models.UserAPI{}, errors.New("Name can not be empty")
	}

	hashedPassword, err := password.Hash(newUser.Password)
	if err != nil {
		return models.UserAPI{}, err
	}

	newUser.Password = hashedPassword
	newUserAPI := models.UserAPI{}

	transactErr := config.Db.Transaction(func (tx *gorm.DB) error {
		userAddRes := tx.Select("name", "email", "password").Create(newUser)

		if userAddRes.Error != nil {
			if strings.HasPrefix(userAddRes.Error.Error(), "Error 1062") {
				return errors.New("Email is already taken")
			}
			return userAddRes.Error
		}

		if userAddRes.RowsAffected == 0 {
			return errors.New("Failed to add user")
		}

		newAdmin := models.Admin{}
		newAdmin.UserID = newUser.UserID
		adminAddRes := tx.Table("admins").Select("user_id").Create(&newAdmin)

		if adminAddRes.Error != nil {
			return adminAddRes.Error
		}

		if adminAddRes.RowsAffected == 0 {
			return errors.New("Failed to add admin")
		}

		return nil

	})
	
	newUserAPI.UserID = newUser.UserID
	newUserAPI.Name = newUser.Name
	newUserAPI.Email = newUser.Email

	if transactErr != nil {
		return models.UserAPI{}, transactErr
	}

	return newUserAPI, nil
}