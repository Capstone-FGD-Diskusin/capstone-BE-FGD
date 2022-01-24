package presentation

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	"github.com/dragranzer/capstone-BE-FGD/features/users"
	"github.com/dragranzer/capstone-BE-FGD/features/users/presentation/request"
	"github.com/dragranzer/capstone-BE-FGD/features/users/presentation/response"
	"github.com/dragranzer/capstone-BE-FGD/middleware"
	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	userBussiness users.Bussiness
}

const (
	projectID  = "triple-bonito-336817" // FILL IN WITH YOURS
	bucketName = "capstone-fgd"         // FILL IN WITH YOURS
)

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

var uploader *ClientUploader

func NewUserHandler(ub users.Bussiness) *UsersHandler {
	return &UsersHandler{
		userBussiness: ub,
	}
}

func (uh *UsersHandler) Register(c echo.Context) error {
	user := request.User{}
	c.Bind(&user)
	err := uh.userBussiness.Register(request.ToCore(user))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "data success di masukkan",
	})
}

func (uH *UsersHandler) LoginUser(c echo.Context) error {
	user := request.User{}
	c.Bind(&user)
	_, token, isAuth, err := uH.userBussiness.Login(request.ToCore(user))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	if !isAuth {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Unauthorized",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Selamat email dan pass mu benar",
		"token":   token,
	})
}

func (uH *UsersHandler) GetProfileData(c echo.Context) error {
	temp := middleware.ExtractClaim(c)
	userID := temp["user_id"].(float64)
	core := users.Core{
		ID: int(userID),
	}
	resp, err := uH.userBussiness.GetProfileData(core)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    response.FromCore(resp),
	})
}

func (uH *UsersHandler) GetUserData(c echo.Context) error {
	var idstring string
	echo.PathParamsBinder(c).String("id", &idstring)
	id, _ := strconv.Atoi(idstring)
	core := users.Core{
		ID: id,
	}
	resp, err := uH.userBussiness.GetProfileData(core)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    response.FromCore(resp),
	})
}

func (uH *UsersHandler) EditUserData(c echo.Context) error {
	user := request.User{}
	c.Bind(&user)
	temp := middleware.ExtractClaim(c)
	userID := temp["user_id"].(float64)
	user.ID = int(userID)
	err := uH.userBussiness.EditDataUser(request.ToCore(user))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func (uH *UsersHandler) DeleteUserDataAdmin(c echo.Context) error {
	var idstring string
	echo.PathParamsBinder(c).String("id", &idstring)
	id, _ := strconv.Atoi(idstring)
	user := users.Core{
		ID: id,
	}
	err := uH.userBussiness.DeleteDataUserbyId(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete data",
	})
}

func (uH *UsersHandler) DeleteUserDataUser(c echo.Context) error {
	user := request.User{}
	temp := middleware.ExtractClaim(c)
	userID := temp["user_id"].(float64)
	user.ID = int(userID)
	err := uH.userBussiness.DeleteDataUserbyId(request.ToCore(user))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func (uH *UsersHandler) UpgradeUserToModerator(c echo.Context) error {
	user := request.UpgradeUser{}
	c.Bind(&user)
	temp := middleware.ExtractClaim(c)
	userID := temp["user_id"].(float64)
	user.AdminID = int(userID)
	err := uH.userBussiness.UpgradeToModerator(request.ToCoreUpgrade(user))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func (uH *UsersHandler) HandleFileUploadToBucket(c echo.Context) error {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "keys.json") // FILL IN WITH YOUR FILE PATH
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	uploader = &ClientUploader{
		cl:         client,
		bucketName: bucketName,
		projectID:  projectID,
		uploadPath: "test-files/",
	}

	f, err := c.FormFile("file_input")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	blobFile, err := f.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err = uploader.UploadFile(blobFile, f.Filename)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success",
		"url":     "https://storage.cloud.google.com/capstone-fgd/test-files/" + f.Filename,
	})
}

func (c *ClientUploader) UploadFile(file multipart.File, object string) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}

func (uH *UsersHandler) GetAllUser(c echo.Context) error {
	user := request.User{}
	c.Bind(&user)
	core := request.ToCore(user)
	resp, err := uH.userBussiness.GetAllUser(core)
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

func (uH *UsersHandler) Ranking(c echo.Context) error {
	resp, err := uH.userBussiness.Ranking()
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
