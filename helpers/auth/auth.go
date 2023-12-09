package auth

import (
	"log"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
)

var TokenAuth = jwtauth.New("HS256", []byte(("TOKEN_SECRET")), nil)

func CreateToken(userId string) (string, error) {
	_, tokenString, _ := TokenAuth.Encode(map[string]interface{}{"userId": userId, "exp": jwtauth.ExpireIn(time.Hour)})
	return tokenString, nil
}

func HashPassword(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func ComparePasswords(hashedPwd, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
