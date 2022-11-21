package users

import (
	"fmt"
	"session-login/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (repository *Repository) Login(user UserLogin) (getUser model.User, err error) {
	query := repository.db.Table("users").Where("username", user.Username).Where("password", user.Password).First(&getUser)

	if query.Error != nil {
		err = fmt.Errorf("error : %s", query.Error)
		return
	}

	return
}

func (repository *Repository) Register(request UserRegister) (getUser model.User, err error) {
	query := repository.db.Table("users").Create(&request)

	if query.Error != nil {
		err = fmt.Errorf("error : %s", query.Error)
		return
	}

	query.Last(&getUser)

	return
}
