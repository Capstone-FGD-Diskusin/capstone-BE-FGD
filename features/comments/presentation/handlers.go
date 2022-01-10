package presentation

import (
	"net/http"

	"github.com/dragranzer/capstone-BE-FGD/features/comments"
	"github.com/dragranzer/capstone-BE-FGD/features/comments/presentation/request"
	"github.com/dragranzer/capstone-BE-FGD/middleware"
	"github.com/labstack/echo/v4"
)

type CommentsHandler struct {
	commentBussiness comments.Bussiness
}

func NewCommentHandler(ub comments.Bussiness) *CommentsHandler {
	return &CommentsHandler{
		commentBussiness: ub,
	}
}

func (ch *CommentsHandler) AddComment(c echo.Context) error {
	comment := request.Comment{}
	c.Bind(&comment)
	temp := middleware.ExtractClaim(c)
	userID := temp["user_id"].(float64)
	comment.UserID = int(userID)
	err := ch.commentBussiness.AddComment(request.ToCore(comment))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "data success di masukkan",
	})
}
