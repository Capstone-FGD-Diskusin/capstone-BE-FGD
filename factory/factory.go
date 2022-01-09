package factory

import (
	"github.com/dragranzer/capstone-BE-FGD/config"
	_user_bussiness "github.com/dragranzer/capstone-BE-FGD/features/users/bussiness"
	_user_data "github.com/dragranzer/capstone-BE-FGD/features/users/data"
	_user_presentation "github.com/dragranzer/capstone-BE-FGD/features/users/presentation"

	_follower_bussiness "github.com/dragranzer/capstone-BE-FGD/features/followers/bussiness"
	_follower_data "github.com/dragranzer/capstone-BE-FGD/features/followers/data"
	_follower_presentation "github.com/dragranzer/capstone-BE-FGD/features/followers/presentation"

	_thread_bussiness "github.com/dragranzer/capstone-BE-FGD/features/threads/bussiness"
	_thread_data "github.com/dragranzer/capstone-BE-FGD/features/threads/data"
	_thread_presentation "github.com/dragranzer/capstone-BE-FGD/features/threads/presentation"

	_like_bussiness "github.com/dragranzer/capstone-BE-FGD/features/likes/bussiness"
	_like_data "github.com/dragranzer/capstone-BE-FGD/features/likes/data"
	_like_presentation "github.com/dragranzer/capstone-BE-FGD/features/likes/presentation"
)

type Presenter struct {
	UserPresentation     *_user_presentation.UsersHandler
	FollowerPresentation *_follower_presentation.FollowersHandler
	ThreadPresentation   *_thread_presentation.ThreadsHandler
	LikePresentation     *_like_presentation.LikesHandler
}

func Init() Presenter {

	userData := _user_data.NewUserRepository(config.DB)
	followerData := _follower_data.NewFollowerRepository(config.DB)
	threadData := _thread_data.NewThreadRepository(config.DB)
	likeData := _like_data.NewLikeRepository(config.DB)

	userBussiness := _user_bussiness.NewUserBussiness(userData)
	followerBussiness := _follower_bussiness.NewFollowerBussiness(followerData)
	threadBussiness := _thread_bussiness.NewThreadBussiness(followerBussiness, threadData)
	likeBussiness := _like_bussiness.NewLikeBussiness(likeData, userBussiness, threadBussiness)

	return Presenter{
		UserPresentation:     _user_presentation.NewUserHandler(userBussiness),
		FollowerPresentation: _follower_presentation.NewFollowerHandler(followerBussiness),
		ThreadPresentation:   _thread_presentation.NewThreadHandler(threadBussiness),
		LikePresentation:     _like_presentation.NewLikeHandler(likeBussiness),
	}
}
