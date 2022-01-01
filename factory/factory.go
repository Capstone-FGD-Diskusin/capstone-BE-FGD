package factory

import (
	"github.com/dragranzer/capstone-BE-FGD/config"
	_user_bussiness "github.com/dragranzer/capstone-BE-FGD/features/users/bussiness"
	_user_data "github.com/dragranzer/capstone-BE-FGD/features/users/data"
	_user_presentation "github.com/dragranzer/capstone-BE-FGD/features/users/presentation"

	_follower_bussiness "github.com/dragranzer/capstone-BE-FGD/features/followers/bussiness"
	_follower_data "github.com/dragranzer/capstone-BE-FGD/features/followers/data"
	_follower_presentation "github.com/dragranzer/capstone-BE-FGD/features/followers/presentation"
)

type Presenter struct {
	UserPresentation     *_user_presentation.UsersHandler
	FollowerPresentation *_follower_presentation.FollowersHandler
}

func Init() Presenter {

	userData := _user_data.NewUserRepository(config.DB)
	followerData := _follower_data.NewFollowerRepository(config.DB)

	userBussiness := _user_bussiness.NewUserBussiness(userData)
	followerBussiness := _follower_bussiness.NewFollowerBussiness(followerData)

	return Presenter{
		UserPresentation:     _user_presentation.NewUserHandler(userBussiness),
		FollowerPresentation: _follower_presentation.NewFollowerHandler(followerBussiness),
	}
}
