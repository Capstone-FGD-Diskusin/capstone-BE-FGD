package presentation

import (
	"net/http"
	"strconv"

	"github.com/dragranzer/capstone-BE-FGD/features/users"
	"github.com/dragranzer/capstone-BE-FGD/features/users/presentation/request"
	"github.com/dragranzer/capstone-BE-FGD/features/users/presentation/response"
	"github.com/dragranzer/capstone-BE-FGD/middleware"
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
	_, token, isAuth, err := uH.userBussiness.Login(request.ToCore(user))
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
		"token":   token,
	})
}

func (uH *UsersHandler) GetProfileData(c echo.Context) error {
	temp := middleware.ExtractClaim(c)
	userID := temp["user_id"].(float64)
	core := users.Core{
		ID: int(userID),
	}
	resp, err := uH.userBussiness.GetProfileData(core)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    response.FromCore(resp),
	})
}

func (uH *UsersHandler) GetUserData(c echo.Context) error {
	var idstring string
	echo.PathParamsBinder(c).String("id", &idstring)
	id, _ := strconv.Atoi(idstring)
	core := users.Core{
		ID: id,
	}
	resp, err := uH.userBussiness.GetProfileData(core)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    response.FromCore(resp),
	})
}

func (uH *UsersHandler) EditUserData(c echo.Context) error {
	user := request.User{}
	c.Bind(&user)
	temp := middleware.ExtractClaim(c)
	userID := temp["user_id"].(float64)
	user.ID = int(userID)
	err := uH.userBussiness.EditDataUser(request.ToCore(user))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
