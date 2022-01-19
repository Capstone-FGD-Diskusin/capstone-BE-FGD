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

	e.GET("/thread/:id/comment", _presenter.CommentPresentation.GetCommentsThread)

	e.POST("/category", _presenter.CategoryPresentation.AddCategory)
	e.PUT("/category/:id", _presenter.CategoryPresentation.EditCategory)
	e.DELETE("/category/:id", _presenter.CategoryPresentation.DeleteCategorybyId)

	eJWT := e.Group("")
	eJWT.Use(mid.JWT([]byte(config.ENV.JWT_SECRET)))

	eJWT.GET("/user", _presenter.UserPresentation.GetProfileData)
	eJWT.PUT("/user", _presenter.UserPresentation.EditUserData)
	eJWT.DELETE("/user", _presenter.UserPresentation.DeleteUserDataUser)

	eJWT.POST("/user/follow", _presenter.FollowerPresentation.Follow)
	eJWT.DELETE("/user/follow", _presenter.FollowerPresentation.Unfollow)
	eJWT.GET("/user/following", _presenter.FollowerPresentation.GetFollowing)

	// eJWT.GET("/thread/homepage", _presenter.ThreadPresentation.GetThreadHome)
	eJWT.POST("/thread", _presenter.ThreadPresentation.AddThread)
	eJWT.GET("/thread/:id", _presenter.ThreadPresentation.GetThread)
	eJWT.DELETE("/thread/:id", _presenter.FavoritePresentation.DeleteThreadbyId)

	eJWT.POST("/like", _presenter.LikePresentation.LikingThread)
	eJWT.DELETE("/like", _presenter.LikePresentation.UnlikingThread)
	eJWT.GET("/thread/homepage", _presenter.LikePresentation.GetThreadHome)

	eJWT.POST("/thread/comment", _presenter.CommentPresentation.AddComment)
	eJWT.DELETE("/thread/comment", _presenter.CommentPresentation.DeleteCommentbyId)

	return e
}
