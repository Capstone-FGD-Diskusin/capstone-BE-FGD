package request

import "github.com/dragranzer/capstone-BE-FGD/features/favorites"

type Favorite struct {
	UserID   int
	ThreadID int `json:"thread_id" form:"thread_id"`
}

func ToCore(req Favorite) favorites.Core {
	return favorites.Core{
		UserID:   req.UserID,
		ThreadID: req.ThreadID,
	}
}
