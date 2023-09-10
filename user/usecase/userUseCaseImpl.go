package usecase

import (
	"gorillaz-clean-v3/models"
	"gorillaz-clean-v3/user"
)

type UserUseCaseImpl struct {
	userRepo user.UserRepo
}

func NewUserUseCase(userRepo user.UserRepo) user.UserUseCase {
	return &UserUseCaseImpl{userRepo}
}

func (r *UserUseCaseImpl) Create(user *models.User) (map[string]interface{}, error) {
	return r.userRepo.Create(user)
}

func (r *UserUseCaseImpl) DeleteId(uid string) error {
	return r.userRepo.DeleteId(uid)
}

func (r *UserUseCaseImpl) GetAll() (*[]models.User, error) {
	return r.userRepo.GetAll()
}

func (r *UserUseCaseImpl) GetSingleId(uid string) (*models.User, error) {
	return r.userRepo.GetSingleId(uid)
}

func (r *UserUseCaseImpl) Login(user *models.User) (map[string]interface{}, error) {
	return r.userRepo.Login(user)
}

func (r *UserUseCaseImpl) UpdateId(uid string) (*models.User, error) {
	return r.userRepo.UpdateId(uid)
}
