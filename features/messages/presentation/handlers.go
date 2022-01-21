package presentation

import (
	"net/http"

	"github.com/dragranzer/capstone-BE-FGD/features/messages"
	"github.com/dragranzer/capstone-BE-FGD/features/messages/presentation/request"
	"github.com/dragranzer/capstone-BE-FGD/middleware"
	"github.com/labstack/echo/v4"
)

type MessagesHandler struct {
	messageBussiness messages.Bussiness
}

func NewMessageHandler(mb messages.Bussiness) *MessagesHandler {
	return &MessagesHandler{
		messageBussiness: mb,
	}
}

func (mh *MessagesHandler) SendMessageToAdmin(c echo.Context) error {
	message := request.Message{}
	c.Bind(&message)
	temp := middleware.ExtractClaim(c)
	ModeratorID := temp["user_id"].(float64)
	message.ModeratorID = int(ModeratorID)
	err := mh.messageBussiness.ReportToAdmin(request.ToCore(message))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "data success di masukkan",
	})
}
