package adapters


import (
	"example/internal/users/domain"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewuserRepository(db *gorm.DB) domain.UserRespository {
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

func (u *userRepo) Update(ID *int, inp *domain.UpdateUserInput) error {
	user, err := u.Get(ID)
	if err != nil {
		return err
	}

	user.UserName = *inp.UserName
	user.Password = *inp.Password

	result := u.db.Save(&user)

	return result.Error
}

func (u *userRepo) FindAll(page, limit int) ([]*domain.User, error) {
	var users []*domain.User

	offset := (page - 1) * limit
	result := u.db.Order("id asc").Limit(limit).Offset(offset).Find(&users)
	return users, result.Error
}

func (u *userRepo) Remove(ID int) error {
	post, err := u.Get(&ID)
	if err != nil {
		return err
	}
	result := u.db.Delete(post)

	return result.Error
}
