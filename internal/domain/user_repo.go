package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName     string `gorm:"type:varchar(255);unique;column:user_name" json:"user_name"`
	FullName     string `gorm:"type:varchar(255);column:full_name" json:"full_name"`
	PasswordHash string `gorm:"type:text;column:password_hash" json:"password"`
}

type GetPaginationInput struct {
	Page  uint `json:"page"`
	Limit uint `json:"limit"`
}

type UserRespository interface {
	Save(user *User) error
	// Get(ID *int) (User, error)

	GetUserName(userName *string) (bool, error)
	GetUser(userName, password_hash string) (*User, error)
	FindAll(page, limit int) ([]*User, error)
	GetUserPassword(password *string) (bool, error)
}
