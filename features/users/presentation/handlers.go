package presentation

import (
	"net/http"

	"github.com/dragranzer/capstone-BE-FGD/features/users"
	"github.com/dragranzer/capstone-BE-FGD/features/users/presentation/request"
	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	userBussiness users.Bussiness
}

func NewUserHandler(ub users.Bussiness) *UsersHandler {
	return &UsersHandler{
		userBussiness: ub,
	}
}

func (uh *UsersHandler) Register(c echo.Context) error {
	user := request.User{}
	c.Bind(&user)
	err := uh.userBussiness.Register(request.ToCore(user))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "data success di masukkan",
	})
}
