package presentation

import (
	"fmt"
	"net/http"

	"github.com/dragranzer/capstone-BE-FGD/features/followers/presentation/request"
	"github.com/dragranzer/capstone-BE-FGD/middleware"
	"github.com/labstack/echo/v4"
)

type FollowersHandler struct {
}

func NewFollowerHandler() *FollowersHandler {
	return &FollowersHandler{}
}

func (fh *FollowersHandler) Follow(c echo.Context) error {

	followed_id := request.Follow{}
	c.Bind(&followed_id)

	temp := middleware.ExtractClaim(c)

	fmt.Println(temp["user_id"])

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "follow berhasil dilakukan",
	})
}
