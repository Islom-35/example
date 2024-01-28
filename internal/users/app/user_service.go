package app

import (
	"example/internal/users/domain"
	
	"time"
)

type UserService interface {
	Create(User *domain.User) error
	Get(ID *int) (domain.User, error)
	Update(ID *int, inp *domain.UpdateUserInput) error
	FindAll(page, limit int) ([]*domain.User, error)
	Remove(ID int) error
}

func NewUserService(repo domain.UserRespository) UserService {
	return &userService{
		repo: repo,
	}
}

type userService struct {
	repo domain.UserRespository
}

func (u *userService)Create(user *domain.User) error {
	user.Created_at = time.Now()
	if err := u.repo.Save(user); err != nil {
		return err
	}
	return nil
}

func (u *userService) Get(ID *int) (domain.User, error) {
	return u.repo.Get(ID)
}

func (u *userService) Update(ID *int, inp *domain.UpdateUserInput) error {
	return u.repo.Update(ID, inp)
}

func (u *userService) FindAll(page, limit int) ([]*domain.User, error) {
	return u.repo.FindAll(page, limit)
}

func (u *userService) Remove(ID int) error {
	return u.repo.Remove(ID)
}
