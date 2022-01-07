package request

import "github.com/dragranzer/capstone-BE-FGD/features/threads"

type UserID struct {
	UserID int `json:"user_id" form:"user_id"`
}

func ToCore(req UserID) threads.Core {
	return threads.Core{
		OwnerID: req.UserID,
	}
}
