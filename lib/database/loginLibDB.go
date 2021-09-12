package libdb

import (
	"altastore/config"
	"altastore/models"
	"altastore/util/jwt"
)

func LoginUser(user *models.User) (string ,error) {
	res := config.Db.Where("email = ? AND password = ?", user.Email, user.Password).First(user)

	if res.Error != nil {
		return "", res.Error
	}

	token, err := implementjwt.CreateToken(int(user.ID))

	if err != nil {
		return "", err
	}

	return token, nil

}