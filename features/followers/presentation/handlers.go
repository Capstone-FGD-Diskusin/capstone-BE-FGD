package presentation

import (
	"net/http"

	"github.com/dragranzer/capstone-BE-FGD/features/followers"
	"github.com/dragranzer/capstone-BE-FGD/features/followers/presentation/request"
	"github.com/dragranzer/capstone-BE-FGD/features/followers/presentation/response"
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
	followingId := temp["user_id"].(float64)

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

func (fh *FollowersHandler) Unfollow(c echo.Context) error {

	followed_id := request.Follow{}
	c.Bind(&followed_id)
	temp := middleware.ExtractClaim(c)
	followingId := temp["user_id"].(float64)

	err := fh.followerBussiness.Unfollow(request.ToCore(followed_id, int(followingId)))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "unfollow berhasil dilakukan",
	})
}

func (fh *FollowersHandler) GetFollowing(c echo.Context) error {
	template := request.Follow{}
	temp := middleware.ExtractClaim(c)
	followingId := temp["user_id"].(float64)
	resp, err := fh.followerBussiness.GetFollowing(request.ToCore(template, int(followingId)))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"IDFollowed": response.FromCoreSlice(resp),
		"message":    "request berhasil",
	})
}

func (fh *FollowersHandler) GetFollowed(c echo.Context) error {
	template := request.Follow{}
	temp := middleware.ExtractClaim(c)
	followedId := temp["user_id"].(float64)
	resp, err := fh.followerBussiness.GetFollowed(request.ToCoreFollowed(template, int(followedId)))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"IDFollowed": response.FromCoreSliceFollowed(resp),
		"message":    "request berhasil",
	})
}
