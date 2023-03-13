package user

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type Repository interface {
	Create(user *User) error
	GetAll() ([]User, error)
	Get(id string) (*User, error)
}

type repo struct {
	log *log.Logger
	db  *gorm.DB
}

func NewRepo(log *log.Logger, db *gorm.DB) Repository {
	return &repo{
		log: log,
		db:  db,
	}
}

func (repo *repo) Create(user *User) error {
	user.ID = uuid.New().String()
	if err := repo.db.Create(user).Error; err != nil {
		fmt.Println(err)
		return err
	}
	repo.log.Println("Usuario creado con id:", user.ID)
	return nil
}
func (repo *repo) GetAll() ([]User, error) {
	var u []User
	result := repo.db.Model(&u).Order("created_at desc").Find(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

func (repo *repo) Get(id string) (*User, error) {
	user := User{ID: id}
	result := repo.db.First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
