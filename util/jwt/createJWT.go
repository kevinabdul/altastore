package implementjwt

import (
	"time"
	"github.com/golang-jwt/jwt"
	"os"
)

func CreateToken(id int) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = id
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}