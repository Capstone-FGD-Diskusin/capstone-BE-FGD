package data

import "github.com/dragranzer/capstone-BE-FGD/features/users"

type User struct {
	ID             int
	Email          string
	Password       string
	Username       string
	Follower       int
	SumLike        int
	SumComment     int
	ProfilePicture string
}

func fromCore(core users.Core) User {
	return User{
		Email:          core.Email,
		Password:       core.Password,
		Username:       core.Username,
		ProfilePicture: core.ProfilePicture,
	}
}
