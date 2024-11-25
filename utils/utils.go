package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 12)
	return string(hash), err
}

func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	signedToken, err := token.SignedString([]byte("secret"))
	return signedToken, err
}

func CheckPassword(hashPassword, Password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(Password))
	return err == nil
}
