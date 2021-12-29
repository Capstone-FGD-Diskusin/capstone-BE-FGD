package factory

import (
	"github.com/dragranzer/capstone-BE-FGD/config"
	_user_bussiness "github.com/dragranzer/capstone-BE-FGD/features/users/bussiness"
	_user_data "github.com/dragranzer/capstone-BE-FGD/features/users/data"
	_user_presentation "github.com/dragranzer/capstone-BE-FGD/features/users/presentation"
)

type Presenter struct {
	UserPresentation *_user_presentation.UsersHandler
}

func Init() Presenter {

	userData := _user_data.NewUserRepository(config.DB)

	userBussiness := _user_bussiness.NewUserBussiness(userData)

	return Presenter{
		UserPresentation: _user_presentation.NewUserHandler(userBussiness),
	}
}
