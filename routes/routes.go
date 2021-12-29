package routes

import (
	"github.com/dragranzer/capstone-BE-FGD/factory"
	"github.com/dragranzer/capstone-BE-FGD/middleware"
	"github.com/labstack/echo/v4"
)

func Setup() *echo.Echo {
	_presenter := factory.Init()
	e := echo.New()

	middleware.LogMidd(e)
	e.POST("/user", _presenter.UserPresentation.Register)

	return e
}
