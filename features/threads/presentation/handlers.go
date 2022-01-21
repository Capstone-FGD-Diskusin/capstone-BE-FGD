package presentation

import (
	"net/http"
	"strconv"

	"github.com/dragranzer/capstone-BE-FGD/features/threads"
	"github.com/dragranzer/capstone-BE-FGD/features/threads/presentation/request"
	"github.com/dragranzer/capstone-BE-FGD/features/threads/presentation/response"
	"github.com/dragranzer/capstone-BE-FGD/middleware"
	"github.com/labstack/echo/v4"
)

type ThreadsHandler struct {
	threadBussiness threads.Bussiness
}

func NewThreadHandler(ub threads.Bussiness) *ThreadsHandler {
	return &ThreadsHandler{
		threadBussiness: ub,
	}
}

// func (uh *ThreadsHandler) GetThreadHome(c echo.Context) error {
// 	req := request.Request{}
// 	c.Bind(&req)
// 	temp := middleware.ExtractClaim(c)
// 	ownerID := temp["user_id"].(float64)
// 	fmt.Println(ownerID)
// 	data := threads.Core{
// 		OwnerID: int(ownerID),
// 		Page:    req.Page,
// 	}
// 	threads, err := uh.threadBussiness.GetThreadHome(data)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"message": err.Error(),
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"threads": response.FromCoreSlice(threads),
// 		"message": "success",
// 	})
// }

func (th *ThreadsHandler) AddThread(c echo.Context) error {
	thread := request.Thread{}
	c.Bind(&thread)
	temp := middleware.ExtractClaim(c)
	userID := temp["user_id"].(float64)
	thread.UserID = int(userID)
	err := th.threadBussiness.AddThread(request.ToCoreThread(thread))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "data success di masukkan",
	})
}

func (th *ThreadsHandler) GetThread(c echo.Context) error {
	var idstring string
	echo.PathParamsBinder(c).String("id", &idstring)
	id, _ := strconv.Atoi(idstring)
	core := threads.Core{
		ID: id,
	}
	resp, err := th.threadBussiness.GetThreadbyID(core)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    response.FromCore(resp),
		"message": "ini dia threadnya :)",
	})
}

func (th *ThreadsHandler) GetThreadAll(c echo.Context) error {
	thread := request.Request{}
	c.Bind(&thread)
	resp, err := th.threadBussiness.GetAllThread(request.ToCore(thread))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    response.FromCoreSlice(resp),
		"message": "ini dia threadnya :)",
	})
}
