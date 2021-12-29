package request

import "github.com/dragranzer/capstone-BE-FGD/features/users"

type User struct {
	Email          string `json:"email" form:"email"`
	Password       string `gorm:"size:100" json:"password" form:"password"`
	Username       string `json:"username" form:"username"`
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
}

func ToCore(req User) users.Core {
	return users.Core{
		Username:       req.Username,
		Email:          req.Email,
		Password:       req.Password,
		ProfilePicture: req.ProfilePicture,
	}
}
