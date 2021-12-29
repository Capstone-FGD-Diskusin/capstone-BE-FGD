package response

import "github.com/dragranzer/capstone-BE-FGD/features/users"

type User struct {
	ID       int
	Username string
	Email    string
}

func FromCore(res users.Core) User {
	return User{
		ID:       res.ID,
		Username: res.Username,
		Email:    res.Email,
	}
}
