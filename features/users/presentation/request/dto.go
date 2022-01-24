package request

import "github.com/dragranzer/capstone-BE-FGD/features/users"

type User struct {
	ID             int
	Email          string `json:"email" form:"email"`
	Password       string `gorm:"size:100" json:"password" form:"password"`
	Username       string `json:"username" form:"username"`
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
	Alamat         string `json:"alamat" form:"alamat"`
	Gender         string `json:"gender" form:"gender"`
	Phone          string `json:"phone" form:"phone"`
	Page           int    `json:"page" form:"page"`
}

type UpgradeUser struct {
	ID         int `json:"id" form:"id"`
	CategoryID int `json:"category_id" form:"category_id"`
	AdminID    int
}

func ToCore(req User) users.Core {
	return users.Core{
		Username:       req.Username,
		Email:          req.Email,
		Password:       req.Password,
		ProfilePicture: req.ProfilePicture,
		Alamat:         req.Alamat,
		Gender:         req.Gender,
		Phone:          req.Phone,
		Page:           req.Page,
	}
}

func ToCoreUpgrade(req UpgradeUser) users.Core {
	return users.Core{
		ID:         req.ID,
		CategoryID: req.CategoryID,
		AdminID:    req.AdminID,
	}
}
