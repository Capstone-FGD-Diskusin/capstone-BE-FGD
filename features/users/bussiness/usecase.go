package bussiness

import "github.com/dragranzer/capstone-BE-FGD/features/users"

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
