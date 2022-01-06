package bussiness

import (
	"fmt"

	"github.com/dragranzer/capstone-BE-FGD/features/followers"
	"github.com/dragranzer/capstone-BE-FGD/features/users"
	"github.com/dragranzer/capstone-BE-FGD/middleware"
	"golang.org/x/crypto/bcrypt"
)

type usersUsecase struct {
	userData         users.Data
	followerBussines followers.Bussiness
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

func (uu *usersUsecase) GetThreadHome(data users.Core) (userData users.Core, err error) {
	idCore := followers.Core{
		FollowingID: data.ID,
	}
	listFollowedID, err := uu.followerBussines.GetFollowing(idCore)
	fmt.Println(listFollowedID)
	return
}
