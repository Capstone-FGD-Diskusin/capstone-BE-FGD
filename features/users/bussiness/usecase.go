package bussiness

import (
	"errors"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/dragranzer/capstone-BE-FGD/config"
	"github.com/dragranzer/capstone-BE-FGD/features/users"
	"github.com/dragranzer/capstone-BE-FGD/middleware"
	"golang.org/x/crypto/bcrypt"
)

type usersUsecase struct {
	userData users.Data
}

func NewUserBussiness(userData users.Data) users.Bussiness {
	return &usersUsecase{
		userData: userData,
	}
}

func (uu *usersUsecase) Register(data users.Core) (err error) {
	err = uu.userData.CreateUser(data)
	return err
}

func (uu *usersUsecase) Login(data users.Core) (userData users.Core, token string, isAuth bool, err error) {

	userData, err = uu.userData.SelectDatabyEmail(data)
	if err != nil {
		return userData, token, false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(data.Password))
	isAuth = (err == nil)

	token, err = middleware.CreateToken(userData.ID, userData.Username)
	if err != nil {
		return userData, token, false, err
	}

	return userData, token, isAuth, err
}

func (uu *usersUsecase) GetProfileData(data users.Core) (resp users.Core, err error) {
	resp, err = uu.userData.SelectDatabyID(data)
	return
}

func (uu *usersUsecase) IncrementLike(data users.Core) (err error) {
	err = uu.userData.UpdateLikebyOne(data)
	return
}

func (uu *usersUsecase) DecrementLike(data users.Core) (err error) {
	err = uu.userData.UpdateMinLikebyOne(data)
	return
}

func (uu *usersUsecase) IncrementThread(data users.Core) (err error) {
	err = uu.userData.UpdateThreadbyOne(data)
	return
}

func (uu *usersUsecase) DecrementThread(data users.Core) (err error) {
	err = uu.userData.UpdateMinThreadbyOne(data)
	return
}

func (uu *usersUsecase) IncrementFol(data users.Core) (err error) {
	err = uu.userData.UpdateFolbyOne(data)
	return
}

func (uu *usersUsecase) DecrementFol(data users.Core) (err error) {
	err = uu.userData.UpdateMinFolbyOne(data)
	return
}

func (uu *usersUsecase) IncrementFollowing(data users.Core) (err error) {
	err = uu.userData.UpdateFollowingbyOne(data)
	return
}

func (uu *usersUsecase) DecrementFollowing(data users.Core) (err error) {
	err = uu.userData.UpdateMinFollowingbyOne(data)
	return
}

func (uu *usersUsecase) EditDataUser(data users.Core) (err error) {
	err = uu.userData.UpdateDataUser(data)
	return
}

func (uu *usersUsecase) DeleteDataUserbyId(data users.Core) (err error) {
	err = uu.userData.DeleteDataUserbyId(data)
	return
}

func (uu *usersUsecase) UpgradeToModerator(data users.Core) (err error) {
	fmt.Println(data)
	adminCore := users.Core{
		ID: data.AdminID,
	}
	admin, err := uu.userData.SelectDatabyID(adminCore)
	if err != nil {
		return
	}
	if admin.Role != "admin" {
		err = errors.New("jadi admin dulu yaaa :v")
	}
	if err != nil {
		return
	}
	err = uu.userData.UpdateUserToModerator(data)
	return
}

func (uu *usersUsecase) UploadImage(data users.Core) (err error) {

	return
}

func (uu *usersUsecase) GetAllUser(data users.Core) (resp []users.Core, err error) {
	resp, err = uu.userData.SelectAllUser(data)
	return
}

func (uu *usersUsecase) Ranking() (resp []users.Core, err error) {
	resp, err = uu.userData.Ranking()
	return
}

func (uu *usersUsecase) SendMail(to []string, subject string, message string) error {
	body := "From: " + config.ENV.CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", config.ENV.CONFIG_AUTH_EMAIL, config.ENV.CONFIG_AUTH_PASSWORD, config.ENV.CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%s", config.ENV.CONFIG_SMTP_HOST, config.ENV.CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, config.ENV.CONFIG_AUTH_EMAIL, append(to), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
