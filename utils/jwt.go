package utils

import (
	"strconv"
	"time"

	"gocoon/config"

	"github.com/golang-jwt/jwt/v5"
)

type jwtUtils struct{}

func Jwt() *jwtUtils {
	jwt := &jwtUtils{}
	return jwt
}

func (jwtUtils) CreateToken(UserID int) (string, error) {
	expiration := config.Data.Jwt.ExpiredAt
	secret := config.Data.Jwt.Secret

	payload := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    strconv.Itoa(UserID),
		"expired_at": time.Duration(time.Now().Unix()) + time.Duration(expiration),
	})

	token, err := payload.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}
