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

func (uu *usersUsecase) Login(data users.Core) (userData users.Core, token string, isAuth bool, err error) {

	userData, err = uu.userData.SelectDatabyEmail(data)
	if err != nil {
		return userData, token, false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(data.Password))
	isAuth = (err == nil)

	token, err = middleware.CreateToken(userData.ID, userData.Username)
	if err != nil {
		return userData, token, false, err
	}

	return userData, token, isAuth, err
}

func (uu *usersUsecase) GetProfileData(data users.Core) (resp users.Core, err error) {
	resp, err = uu.userData.SelectDatabyID(data)
	return
}

func (uu *usersUsecase) IncrementLike(data users.Core) (err error) {
	err = uu.userData.UpdateLikebyOne(data)
	return
}
