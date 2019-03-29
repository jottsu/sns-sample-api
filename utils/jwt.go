package utils

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jottsu/sns-sample-api/models"
)

func CreateUserJwt(user *models.User) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_TOKEN")))
	return tokenString, err
}
