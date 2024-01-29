package domain

import (
	"time"
)

type User struct {
	Id         int    `gorm:"primarykey" json:"id"`
	UserName   string `json:"user_name"`
	FullName   string `json:"full_name"`
	Password   string `json:"password"`
	Created_at time.Time
}

type GetPaginationInput struct {
	Page  uint `json:"page"`
	Limit uint `json:"limit"`
}

type UpdateProductInput struct {
	Name  *string `json:"name"`
	Price *int    `json:"price"`
}

type UserRespository interface {
	Save(user *User) error
	Get(ID *int) (User, error)
	GetPassword(password *string) (bool, error)
	GetFullName(fullName *string) (bool, error)
	FindAll(page, limit int) ([]*User, error)
}

