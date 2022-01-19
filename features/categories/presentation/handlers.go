package presentation

import (
	"net/http"

	"github.com/dragranzer/capstone-BE-FGD/features/categories"
	"github.com/dragranzer/capstone-BE-FGD/features/categories/presentation/request"
	"github.com/labstack/echo/v4"
)

type CategorysHandler struct {
	categoryBussiness categories.Bussiness
}

func NewCategoryHandler(cb categories.Bussiness) *CategorysHandler {
	return &CategorysHandler{
		categoryBussiness: cb,
	}
}

func (ch *CategorysHandler) AddCategory(c echo.Context) error {
	category := request.Category{}
	c.Bind(&category)
	err := ch.categoryBussiness.AddCategory(request.ToCore(category))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "data success di masukkan",
	})
}
