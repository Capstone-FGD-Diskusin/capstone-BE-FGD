package routes

import (
	"github.com/dragranzer/capstone-BE-FGD/config"
	"github.com/dragranzer/capstone-BE-FGD/factory"
	"github.com/dragranzer/capstone-BE-FGD/middleware"
	"github.com/labstack/echo/v4"

	mid "github.com/labstack/echo/v4/middleware"
)

func Setup() *echo.Echo {
	_presenter := factory.Init()
	e := echo.New()

	middleware.LogMidd(e)
	middleware.CorsAuth(e)
	e.POST("/user", _presenter.UserPresentation.Register)
	e.POST("/user/login", _presenter.UserPresentation.LoginUser)
	e.GET("/user/:id", _presenter.UserPresentation.GetUserData)
	e.DELETE("/user/:id", _presenter.UserPresentation.DeleteUserDataAdmin)
	e.GET("/user/:id/threads", _presenter.ThreadPresentation.GetThreadUser)
	e.POST("/user/picture", _presenter.UserPresentation.HandleFileUploadToBucket)
	e.GET("/all_user", _presenter.UserPresentation.GetAllUser)
	e.GET("/user/ranking", _presenter.UserPresentation.Ranking)
	e.POST("/user/forget", _presenter.UserPresentation.ForgetPassword)

	e.GET("/thread/:id/comment", _presenter.CommentPresentation.GetCommentsThread)
	e.GET("/thread/comment/:id/balasan", _presenter.CommentPresentation.GetBalasanCommentbyId)
	e.GET("/thread/search", _presenter.CommentPresentation.SearchThread)
	e.GET("/thread", _presenter.ThreadPresentation.GetThreadAll)

	e.POST("/category", _presenter.CategoryPresentation.AddCategory)
	e.PUT("/category/:id", _presenter.CategoryPresentation.EditCategory)
	e.DELETE("/category/:id", _presenter.CategoryPresentation.DeleteCategorybyId)
	e.GET("/category", _presenter.CategoryPresentation.GetAllCategory)

	eJWT := e.Group("")
	eJWT.Use(mid.JWT([]byte(config.ENV.JWT_SECRET)))

	eJWT.GET("/user", _presenter.UserPresentation.GetProfileData)
	eJWT.PUT("/user", _presenter.UserPresentation.EditUserData)
	eJWT.DELETE("/user", _presenter.UserPresentation.DeleteUserDataUser)

	eJWT.POST("/user/follow", _presenter.FollowerPresentation.Follow)
	eJWT.DELETE("/user/follow", _presenter.FollowerPresentation.Unfollow)
	eJWT.DELETE("/user/followed", _presenter.FollowerPresentation.PaksaUnfollow)
	eJWT.GET("/user/following", _presenter.FollowerPresentation.GetFollowing)
	eJWT.GET("/user/followed", _presenter.FollowerPresentation.GetFollowed)

	eJWT.GET("/user/favorite", _presenter.FavoritePresentation.GetAllfavoriteUser)

	// eJWT.GET("/thread/homepage", _presenter.ThreadPresentation.GetThreadHome)
	eJWT.POST("/thread", _presenter.ThreadPresentation.AddThread)
	eJWT.GET("/thread/:id", _presenter.ThreadPresentation.GetThread)
	eJWT.DELETE("/thread/:id", _presenter.FavoritePresentation.DeleteThreadbyId)

	eJWT.POST("/like", _presenter.LikePresentation.LikingThread)
	eJWT.DELETE("/like", _presenter.LikePresentation.UnlikingThread)
	eJWT.GET("/thread/homepage", _presenter.LikePresentation.GetThreadHome)

	eJWT.POST("/thread/comment", _presenter.CommentPresentation.AddComment)
	eJWT.DELETE("/thread/comment", _presenter.CommentPresentation.DeleteCommentbyId)

	eJWT.POST("/message", _presenter.MessagePresentation.SendMessageToAdmin)
	eJWT.GET("/message/admin", _presenter.MessagePresentation.GetMessagebyAdminID)
	eJWT.DELETE("/message/:id", _presenter.MessagePresentation.DeleteMessagebyId)

	eJWT.PUT("/user/upgrade", _presenter.UserPresentation.UpgradeUserToModerator)

	eJWT.POST("/user/thread/follow", _presenter.FavoritePresentation.Insertfavorite)
	eJWT.DELETE("/user/thread/follow", _presenter.FavoritePresentation.Deletefavorite)

	return e
}
