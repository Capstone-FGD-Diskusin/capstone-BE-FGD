package presentation

import (
	"net/http"
	"strconv"

	"github.com/dragranzer/capstone-BE-FGD/features/favorites"
	"github.com/dragranzer/capstone-BE-FGD/features/favorites/presentation/request"
	"github.com/dragranzer/capstone-BE-FGD/features/favorites/presentation/response"
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

func (fh *FavoritesHandler) Insertfavorite(c echo.Context) error {
	req := request.Favorite{}
	c.Bind(&req)
	temp := middleware.ExtractClaim(c)
	userID := temp["user_id"].(float64)
	req.UserID = int(userID)

	err := fh.favoriteBussiness.InsertFavorite(request.ToCore(req))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func (fh *FavoritesHandler) Deletefavorite(c echo.Context) error {
	req := request.Favorite{}
	c.Bind(&req)
	temp := middleware.ExtractClaim(c)
	userID := temp["user_id"].(float64)
	req.UserID = int(userID)

	err := fh.favoriteBussiness.DeleteFavorite(request.ToCore(req))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func (fh *FavoritesHandler) GetAllfavoriteUser(c echo.Context) error {
	req := request.Favorite{}
	c.Bind(&req)
	temp := middleware.ExtractClaim(c)
	userID := temp["user_id"].(float64)
	req.UserID = int(userID)

	resp, err := fh.favoriteBussiness.GetAllFavorite(request.ToCore(req))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    response.FromCoreSlice(resp),
	})
}
