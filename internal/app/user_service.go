package app

import (
	"crypto/sha1"
	"errors"
	"example/internal/domain"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = ""
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
type UserService interface {
	SignUp(user domain.User) error
	LoginUser(userName, pass string) error
	FindAll(page, limit int) ([]*domain.User, error)
	ParseToken(accessToken string) (int, error)
	GenerateToken(username, password string) (string, error)
	// generatePasswordHash(password string) string
}

func NewUserService(repo domain.UserRespository) UserService {
	return &userService{
		repo: repo,
	}
}

type userService struct {
	repo domain.UserRespository
}

func (u *userService) LoginUser(userName, pass string) error {
	log.Println(pass)
	_, err := u.repo.GetUser(userName, pass)
	if err != nil {
		return domain.ErrUserNotFound
	}

	return nil
}

func (u *userService) SignUp(user domain.User) error {
	err := Checker(user)
	if err != nil {
		return err
	}

	user.PasswordHash = generatePasswordHash(user.PasswordHash)

	_, err = u.repo.GetUserName(&user.UserName)
	if err != nil {

		return domain.ErrUserAlreadyExists
	}

	if err := u.repo.Save(&user); err != nil {
		return err
	}
	return nil
}

func (u *userService) FindAll(page, limit int) ([]*domain.User, error) {
	return u.repo.FindAll(page, limit)
}

func Checker(user domain.User) error {
	if user.PasswordHash == "" {
		return domain.ErrInvalidPassword
	}
	if user.FullName == "" {
		return domain.ErrInvalidFullName
	}
	if user.UserName == "" {
		return domain.ErrInvalidUserName
	}
	return nil
}

func (u *userService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (u *userService) GenerateToken(username, password string) (string, error) {
	genPassword := generatePasswordHash(password)
	user, err := u.repo.GetUser(username, genPassword)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		int(user.ID),
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
