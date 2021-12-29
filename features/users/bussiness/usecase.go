package bussiness

import (
	"github.com/dragranzer/capstone-BE-FGD/features/users"
	"github.com/dragranzer/capstone-BE-FGD/middleware"
	"golang.org/x/crypto/bcrypt"
)

type usersUsecase struct {
	userData users.Data
}

func NewUserBussiness(userData users.Data) users.Bussiness {
	return &usersUsecase{
		userData: userData,
	}
}

func (uu *usersUsecase) Register(data users.Core) (err error) {
	err = uu.userData.CreateUser(data)
	return err
}

func (uu *usersUsecase) Login(email string, pass string) (userData users.Core, token string, isAuth bool, err error) {

	userData, err = uu.userData.SelectDatabyEmail(email)
	if err != nil {
		return userData, token, false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(pass))
	isAuth = (err == nil)

	token, err = middleware.CreateToken(userData.ID, userData.Username)
	if err != nil {
		return userData, token, false, err
	}

	return userData, token, isAuth, err
}
