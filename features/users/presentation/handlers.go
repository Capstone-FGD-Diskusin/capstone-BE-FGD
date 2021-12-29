package presentation

import (
	"net/http"

	"github.com/dragranzer/capstone-BE-FGD/features/users"
	"github.com/dragranzer/capstone-BE-FGD/features/users/presentation/request"
	"github.com/dragranzer/capstone-BE-FGD/features/users/presentation/response"
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

func (uH *UsersHandler) LoginUser(c echo.Context) error {
	user := request.User{}
	c.Bind(&user)
	resp, token, isAuth, err := uH.userBussiness.Login(user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	if !isAuth {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Unauthorized",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Selamat email dan pass mu benar",
		"data":    response.FromCore(resp),
		"token":   token,
	})
}
