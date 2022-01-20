package response

import "github.com/dragranzer/capstone-BE-FGD/features/users"

type User struct {
	ID             int
	Username       string
	Email          string
	Alamat         string
	Gender         string
	Phone          string
	Follower       int
	SumLike        int
	SumComment     int
	ProfilePicture string
}

func FromCore(res users.Core) User {
	return User{
		ID:             res.ID,
		Username:       res.Username,
		Email:          res.Email,
		Follower:       res.Follower,
		SumLike:        res.SumLike,
		SumComment:     res.SumComment,
		ProfilePicture: res.ProfilePicture,
		Alamat:         res.Alamat,
		Gender:         res.Gender,
		Phone:          res.Phone,
	}
}
