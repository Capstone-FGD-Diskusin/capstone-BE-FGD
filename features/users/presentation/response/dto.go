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
	Following      int
	SumLike        int
	SumComment     int
	ProfilePicture string
	Role           string
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
		Following:      res.Following,
		Role:           res.Role,
	}
}

func FromCoreSlice(data []users.Core) []User {
	resp := []User{}
	for _, value := range data {
		resp = append(resp, FromCore(value))
	}
	return resp
}
