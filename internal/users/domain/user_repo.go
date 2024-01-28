package domain

import (
	"errors"
	"time"
)

var (
	ErrUserNotFound = errors.New("UserNotFound")
)

type User struct {
	Id         int    `gorm:"primarykey" json:"id"`
	UserName       string `json:"name"`
	FullName       string `json:"full_name"`
	Password	string `json:"password"`
	Created_at time.Time
}

type GetPaginationInput struct {
	Page  uint `json:"page"`
	Limit uint `json:"limit"`
}

type UpdateUserInput struct {
	UserName  *string `json:"name"`
	Password *string    `json:"price"`
}


type UserRespository interface {
	Save(user *User) error
	Get(ID *int) (User, error)
	Update(ID *int, inp *UpdateUserInput) error
	FindAll(page, limit int) ([]*User, error)
	Remove(ID int) error
}
