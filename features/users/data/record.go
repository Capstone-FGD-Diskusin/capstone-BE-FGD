package data

import (
	"time"

	"github.com/dragranzer/capstone-BE-FGD/features/users"
)

type User struct {
	ID             int
	Email          string
	Password       string
	Username       string
	Alamat         string
	Gender         string
	Phone          string
	Follower       int
	Following      int
	SumLike        int
	SumComment     int
	SumThread      int
	ProfilePicture string
	Role           string
	CategoryID     int
	CreatedAt      time.Time
}

func fromCore(core users.Core) User {
	return User{
		Email:          core.Email,
		Password:       core.Password,
		Username:       core.Username,
		ProfilePicture: core.ProfilePicture,
		Alamat:         core.Alamat,
		Gender:         core.Gender,
		Phone:          core.Phone,
		CategoryID:     core.CategoryID,
		SumLike:        core.SumLike,
		SumComment:     core.SumComment,
		SumThread:      core.SumThread,
	}
}

func (a *User) toCore() users.Core {
	return users.Core{
		ID:             int(a.ID),
		Username:       a.Username,
		Email:          a.Email,
		Password:       a.Password,
		Follower:       a.Follower,
		ProfilePicture: a.ProfilePicture,
		SumLike:        a.SumLike,
		SumComment:     a.SumComment,
		Role:           a.Role,
		CategoryID:     a.CategoryID,
		Gender:         a.Gender,
		Phone:          a.Phone,
		Alamat:         a.Alamat,
		Following:      a.Following,
		SumThread:      a.SumThread,
	}
}

func ToCoreSlice(data []User) []users.Core {
	resp := []users.Core{}
	for _, value := range data {
		resp = append(resp, value.toCore())
	}
	return resp
}
