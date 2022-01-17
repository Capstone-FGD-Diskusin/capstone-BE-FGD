package presentation

import (
	"net/http"
	"strconv"

	"github.com/dragranzer/capstone-BE-FGD/features/favorites"
	"github.com/dragranzer/capstone-BE-FGD/features/favorites/presentation/request"
	"github.com/dragranzer/capstone-BE-FGD/middleware"
	"github.com/labstack/echo/v4"
)

type FavoritesHandler struct {
	favoriteBussiness favorites.Bussiness
}

func NewFavoriteHandler(ub favorites.Bussiness) *FavoritesHandler {
	return &FavoritesHandler{
		favoriteBussiness: ub,
	}
}

func (fh *FavoritesHandler) DeleteThreadbyId(c echo.Context) error {
	var idstring string
	echo.PathParamsBinder(c).String("id", &idstring)
	id_thread, _ := strconv.Atoi(idstring)

	favorite := request.Favorite{}
	temp := middleware.ExtractClaim(c)
	userID := temp["user_id"].(float64)

	favorite.UserID = int(userID)
	favorite.ThreadID = id_thread

	err := fh.favoriteBussiness.DeleteThreadbyId(request.ToCore(favorite))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "data success di hapus",
	})
}
