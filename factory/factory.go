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

	_comment_bussiness "github.com/dragranzer/capstone-BE-FGD/features/comments/bussiness"
	_comment_data "github.com/dragranzer/capstone-BE-FGD/features/comments/data"
	_comment_presentation "github.com/dragranzer/capstone-BE-FGD/features/comments/presentation"

	_favorite_bussiness "github.com/dragranzer/capstone-BE-FGD/features/favorites/bussiness"
	_favorite_data "github.com/dragranzer/capstone-BE-FGD/features/favorites/data"
	_favorite_presentation "github.com/dragranzer/capstone-BE-FGD/features/favorites/presentation"

	_category_bussiness "github.com/dragranzer/capstone-BE-FGD/features/categories/bussiness"
	_category_data "github.com/dragranzer/capstone-BE-FGD/features/categories/data"
	_category_presentation "github.com/dragranzer/capstone-BE-FGD/features/categories/presentation"

	_message_bussiness "github.com/dragranzer/capstone-BE-FGD/features/messages/bussiness"
	_message_data "github.com/dragranzer/capstone-BE-FGD/features/messages/data"
	_message_presentation "github.com/dragranzer/capstone-BE-FGD/features/messages/presentation"
)

type Presenter struct {
	UserPresentation     *_user_presentation.UsersHandler
	FollowerPresentation *_follower_presentation.FollowersHandler
	ThreadPresentation   *_thread_presentation.ThreadsHandler
	LikePresentation     *_like_presentation.LikesHandler
	CommentPresentation  *_comment_presentation.CommentsHandler
	FavoritePresentation *_favorite_presentation.FavoritesHandler
	CategoryPresentation *_category_presentation.CategorysHandler
	MessagePresentation  *_message_presentation.MessagesHandler
}

func Init() Presenter {

	userData := _user_data.NewUserRepository(config.DB)
	followerData := _follower_data.NewFollowerRepository(config.DB)
	threadData := _thread_data.NewThreadRepository(config.DB)
	likeData := _like_data.NewLikeRepository(config.DB)
	commentData := _comment_data.NewCommentRepository(config.DB)
	favoriteData := _favorite_data.NewFavoriteRepository(config.DB)
	categoryData := _category_data.NewCategoryRepository(config.DB)
	messageData := _message_data.NewMessageRepository(config.DB)

	userBussiness := _user_bussiness.NewUserBussiness(userData)
	categoryBussiness := _category_bussiness.NewCategoryBussiness(categoryData)
	followerBussiness := _follower_bussiness.NewFollowerBussiness(followerData, userBussiness)
	threadBussiness := _thread_bussiness.NewThreadBussiness(followerBussiness, threadData, categoryBussiness)
	likeBussiness := _like_bussiness.NewLikeBussiness(likeData, userBussiness, threadBussiness)
	commentBussiness := _comment_bussiness.NewCommentBussiness(commentData, threadBussiness)
	favoriteBussiness := _favorite_bussiness.NewFavoriteBussiness(threadBussiness, userBussiness, commentBussiness, favoriteData, likeBussiness)
	messageBussiness := _message_bussiness.NewMessageBussiness(messageData, userBussiness, threadBussiness, commentBussiness)

	return Presenter{
		UserPresentation:     _user_presentation.NewUserHandler(userBussiness),
		FollowerPresentation: _follower_presentation.NewFollowerHandler(followerBussiness),
		ThreadPresentation:   _thread_presentation.NewThreadHandler(threadBussiness),
		LikePresentation:     _like_presentation.NewLikeHandler(likeBussiness),
		CommentPresentation:  _comment_presentation.NewCommentHandler(commentBussiness),
		FavoritePresentation: _favorite_presentation.NewFavoriteHandler(favoriteBussiness),
		CategoryPresentation: _category_presentation.NewCategoryHandler(categoryBussiness),
		MessagePresentation:  _message_presentation.NewMessageHandler(messageBussiness),
	}
}
