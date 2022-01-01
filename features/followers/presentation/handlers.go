package presentation

import (
	"net/http"

	"github.com/dragranzer/capstone-BE-FGD/features/followers"
	"github.com/dragranzer/capstone-BE-FGD/features/followers/presentation/request"
	"github.com/dragranzer/capstone-BE-FGD/middleware"
	"github.com/labstack/echo/v4"
)

type FollowersHandler struct {
	followerBussiness followers.Bussiness
}

func NewFollowerHandler(fb followers.Bussiness) *FollowersHandler {
	return &FollowersHandler{
		followerBussiness: fb,
	}
}

func (fh *FollowersHandler) Follow(c echo.Context) error {

	followed_id := request.Follow{}
	c.Bind(&followed_id)

	temp := middleware.ExtractClaim(c)

	// fmt.Println(temp["user_id"])

	followingId := temp["user_id"].(float64)

	// fmt.Println(followingId)

	err := fh.followerBussiness.Follow(request.ToCore(followed_id, int(followingId)))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "follow berhasil dilakukan",
	})
}
