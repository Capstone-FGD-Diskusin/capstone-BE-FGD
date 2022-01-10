package request

import "github.com/dragranzer/capstone-BE-FGD/features/likes"

type Thread struct {
	ThreadID int `json:"thread_id" form:"thread_id"`
}

type Request struct {
	Page int `json:"page" form:"page"`
}

func ToCore(req Thread, userId int) likes.Core {

	return likes.Core{
		UserID:   userId,
		ThreadID: req.ThreadID,
	}
}
