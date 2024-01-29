package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model       
	UserName string `json:"user_name"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

type GetPaginationInput struct {
	Page  uint `json:"page"`
	Limit uint `json:"limit"`
}

type UserRespository interface {
	Save(user *User) error
	Get(ID *int) (User, error)
	GetPassword(password *string) (bool, error)
	GetUserName(userName *string) (bool, error)
	FindAll(page, limit int) ([]*User, error)
}
