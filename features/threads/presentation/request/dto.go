package request

import "github.com/dragranzer/capstone-BE-FGD/features/threads"

type Request struct {
	UserID int `json:"user_id" form:"user_id"`
	Page   int `json:"page" form:"page"`
}

func ToCore(req Request) threads.Core {
	return threads.Core{
		OwnerID: req.UserID,
		Page:    req.Page,
	}
}
