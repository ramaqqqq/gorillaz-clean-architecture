package user

import "gorillaz-clean-v3/models"

// type R map[string]interface{}

type UserRepo interface {
	Login(user *models.User) (map[string]interface{}, error)
	Create(user *models.User) (map[string]interface{}, error)
	GetAll() (*[]models.User, error)
	GetSingleId(uid string) (*models.User, error)
	UpdateId(uid string) (*models.User, error)
	DeleteId(uid string) error
}
