package repo

import (
	"errors"
	"gorillaz-clean-v3/helpers"
	"gorillaz-clean-v3/middleware"
	"gorillaz-clean-v3/models"
	"gorillaz-clean-v3/user"
	"strconv"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type R map[string]interface{}

type UserRepoImpl struct {
	DB *gorm.DB
}

func NewUserRepo(DB *gorm.DB) user.UserRepo {
	return &UserRepoImpl{DB}
}

func (r *UserRepoImpl) Create(user *models.User) (map[string]interface{}, error) {

	hashedPassword, err := helpers.Hash(user.Password)
	if err != nil {
		helpers.Logger("error", "In Server: repo user, "+err.Error())
		return nil, err
	}

	datum := models.User{}
	datum.Id = user.Id
	datum.Username = user.Username
	datum.Email = user.Email
	datum.Role = user.Role
	datum.Password = string(hashedPassword)

	err = r.DB.Debug().Create(&user).Error
	if err != nil {
		helpers.Logger("error", "In Server: repo user, "+err.Error())
		return nil, err
	}

	rMsg := R{}
	rMsg["user_id"] = user.Id
	rMsg["username"] = user.Username
	rMsg["email"] = user.Email
	rMsg["role"] = true

	return map[string]interface{}{"users": rMsg}, nil
}

func (r *UserRepoImpl) DeleteId(uid string) error {
	var user = models.User{}
	err := r.DB.Debug().Table("users").Where("id = ?", uid).First(&user).Delete(&user).Error
	if err != nil {
		helpers.Logger("error", "In Server: repo user, "+err.Error())
		return nil
	}
	return nil
}

func (r *UserRepoImpl) GetAll() (*[]models.User, error) {
	var datum []models.User
	err := r.DB.Debug().Find(&datum).Error
	if err != nil {
		helpers.Logger("error", "In Server: repo user, "+err.Error())
		return nil, err
	}
	return &datum, nil
}

func (r *UserRepoImpl) GetSingleId(uid string) (*models.User, error) {
	var datum = models.User{}
	err := r.DB.Debug().Table("users").Where("id = ?", uid).First(&datum).Error
	if err != nil {
		helpers.Logger("error", "In Server: repo user, "+err.Error())
		return nil, err
	}
	return &datum, nil
}

func (r *UserRepoImpl) Login(user *models.User) (map[string]interface{}, error) {
	datum := models.User{}

	err := r.DB.Debug().Where("email = ?", user.Email).Take(&datum).Error
	if err != nil {
		helpers.Logger("error", "In Server: repo user, "+err.Error())
		return nil, err
	}

	err = helpers.VerifyPassword(datum.Password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		helpers.Logger("error", "In Server: invalid verify"+err.Error())
		return nil, errors.New("Password is Failure")
	}

	idInt := strconv.Itoa(datum.Id)
	token, err := middleware.CreateToken(idInt, datum.Username, datum.Email, datum.Role)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	rMsg := R{}
	rMsg["user_id"] = datum.Id
	rMsg["username"] = datum.Username
	rMsg["email"] = datum.Email
	rMsg["role"] = datum.Role
	access := token["accessToken"]
	refresh := token["refreshToken"]

	return R{"accessToken": access, "refreshToken": refresh, "users": rMsg}, nil
}

func (r *UserRepoImpl) UpdateId(uid string) (*models.User, error) {
	var datum = models.User{}
	err := r.DB.Debug().Table("users").Where("id = ?", uid).First(&datum).Update(&datum).Error
	if err != nil {
		helpers.Logger("error", "In Server: repo user, "+err.Error())
		return nil, err
	}
	return &datum, err
}
