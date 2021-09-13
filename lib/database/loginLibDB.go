package libdb

import (
	"altastore/config"
	"altastore/models"
	"altastore/util/jwt"
	"altastore/util/password"
	"errors"
)

func LoginUser(user *models.User) (string ,error) {
	targetUser := models.User{}
	res := config.Db.Where("email = ?", user.Email).First(&targetUser)

	if res.RowsAffected == 0 {
		return "", errors.New("No user with corresponding email")
	}
	
	if res.Error != nil {
		return "", res.Error
	}

	if _, err := password.Check(targetUser.Password, user.Password); err != nil {
		if err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
			return "", errors.New("Given password is incorrect")
		}
		return "", err
	}

	token, err := implementjwt.CreateToken(int(targetUser.ID))

	if err != nil {
		return "", err
	}

	return token, nil

}