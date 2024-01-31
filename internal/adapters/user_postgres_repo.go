package adapters

import (
	"errors"
	"example/internal/domain"
	"log"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRespository {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Save(user *domain.User) error {
	err := u.db.Create(&user)
	return err.Error
}

func (u *userRepo) Get(ID *int) (domain.User, error) {
	var user *domain.User
	result := u.db.First(&user, &ID)

	return *user, result.Error
}

func (u *userRepo) GetUser(userName, passwordHash string) (*domain.User, error) {
	var user *domain.User
	result := u.db.Where(&domain.User{UserName: userName, PasswordHash: passwordHash}).First(&user)
	log.Println(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil 
		}
		return nil, result.Error
	}

	return user, nil
}

func (u *userRepo) GetUserName(userName *string) (bool, error) {
	ok := true

	result := u.db.Where("full_name = ?", *userName)

	if result.Error != nil {
		return false, result.Error
	}

	return ok, nil
}

func (u *userRepo) GetUserPassword(password *string) (bool, error) {
	ok := true

	result := u.db.Where("password_hash = ?", *password)

	if result.Error != nil {
		return false, result.Error
	}

	return ok, nil
}

func (u *userRepo) FindAll(page, limit int) ([]*domain.User, error) {
	var users []*domain.User

	offset := (page - 1) * limit
	result := u.db.Order("id asc").Limit(limit).Offset(offset).Find(&users)
	return users, result.Error
}
