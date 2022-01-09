package presentation

import (
	"net/http"

	"github.com/dragranzer/capstone-BE-FGD/features/likes"
	"github.com/dragranzer/capstone-BE-FGD/features/likes/presentation/request"
	"github.com/dragranzer/capstone-BE-FGD/middleware"
	"github.com/labstack/echo/v4"
)

type LikesHandler struct {
	likeBussiness likes.Bussiness
}

func NewLikeHandler(lb likes.Bussiness) *LikesHandler {
	return &LikesHandler{
		likeBussiness: lb,
	}
}

func (lh *LikesHandler) LikingThread(c echo.Context) error {

	thread_id := request.Thread{}
	c.Bind(&thread_id)
	temp := middleware.ExtractClaim(c)
	followingId := temp["user_id"].(float64)

	err := lh.likeBussiness.LikingThread(request.ToCore(thread_id, int(followingId)))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "follow berhasil dilakukan",
	})
}
