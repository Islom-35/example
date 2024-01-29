package jwt

import (
	"example/pkg/config"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(sub string) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(15 * time.Minute).Unix(),
		"iat": time.Now().Unix(),
		"sub": sub,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecretKey := config.NewConfig().JwtSecretKey
	token, err := jwtToken.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}
