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
	}
}
