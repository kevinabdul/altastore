package libdb

import (
	"altastore/config"
	"altastore/models"
	"altastore/util/password"
	"errors"
	"strings"
	//"fmt"
)

func GetUsers(tableName string) ([]models.UserAPI, error) {
	if tableName == "" {
		tableName = "users"
	}
	var users []models.UserAPI

	res := config.Db.Table(tableName).Find(&users)

	if res.Error != nil {
		if strings.HasPrefix(res.Error.Error(), "Error 1146") {
			return nil, errors.New("Table doesnt exist")
		}
		return nil, res.Error
	}
	return users, nil
}


func GetUserById(targetId int) (models.UserAPI,  error) {
	var user models.UserAPI

	res := config.Db.Model(&models.User{}).Find(&user, targetId)
	
	if res.Error != nil {
		return models.UserAPI{},res.Error
	}

	if res.RowsAffected == 0 {
		return models.UserAPI{}, errors.New("Wrong User Id")
	}
	
	return user, nil
}

func AddUser(newUser *models.User) (models.UserAPI, error) {
	if newUser.Email == "" || newUser.Password == "" {
		return models.UserAPI{}, errors.New("Invalid Email or Password. Make sure its not empty and are of string type")
	}

	if newUser.Name == "" {
		return models.UserAPI{}, errors.New("Name cant be empty")
	}

	hashedPassword, err := password.Hash(newUser.Password)
	if err != nil {
		return models.UserAPI{}, err
	}
	newUser.Password = hashedPassword

	res := config.Db.Select("name", "email", "password").Create(newUser)
	if res.Error != nil {
		return models.UserAPI{}, res.Error
	}
	newUserAPI := models.UserAPI{}
	newUserAPI.UserID = newUser.UserID
	newUserAPI.Name = newUser.Name
	newUserAPI.Email = newUser.Email
	
	return newUserAPI, nil
}

func EditUser(newData models.User, targetId int) (models.UserAPI, error) {
	targetUser := models.User{}

	res := config.Db.Where(`user_id = ?`, targetId).Find(&targetUser).Omit("password", "id").Updates(newData)

	if res.Error != nil {
		return models.UserAPI{},res.Error
	}

	if res.RowsAffected == 0 {
		return models.UserAPI{}, errors.New("Wrong User Id")
	}

	edittedUser := models.UserAPI{}
	edittedUser.UserID = targetUser.UserID
	edittedUser.Name = targetUser.Name
	edittedUser.Email = targetUser.Email

	return edittedUser, nil
}

func DeleteUser(targetId int) (error) {	
	targetUser := models.User{}
	res := config.Db.Find(&targetUser, targetId)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("Wrong User Id")
	}

	config.Db.Exec("set foreign_key_checks = 0")

	res = config.Db.Unscoped().Delete(&targetUser)

	config.Db.Exec("set foreign_key_checks = 1")

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("Failed to delete user")
	}

	return nil
}