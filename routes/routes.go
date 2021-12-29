package routes

import (
	"github.com/dragranzer/capstone-BE-FGD/factory"
	"github.com/labstack/echo/v4"
)

func Setup() *echo.Echo {
	_presenter := factory.Init()
	e := echo.New()

	e.POST("/user", _presenter.UserPresentation.Register)

	return e
}
