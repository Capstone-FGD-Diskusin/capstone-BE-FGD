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
	e.POST("/user", _presenter.UserPresentation.Register)
	e.POST("/user/login", _presenter.UserPresentation.LoginUser)

	eJWT := e.Group("")
	eJWT.Use(mid.JWT([]byte(config.ENV.JWT_SECRET)))

	eJWT.POST("/user/follow", _presenter.FollowerPresentation.Follow)
	eJWT.POST("/user/unfollow", _presenter.FollowerPresentation.Unfollow)
	eJWT.GET("/user/following", _presenter.FollowerPresentation.GetFollowing)

	eJWT.GET("/thread/homepage", _presenter.ThreadPresentation.GetThreadHome)

	return e
}
