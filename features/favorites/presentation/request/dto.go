package request

import "github.com/dragranzer/capstone-BE-FGD/features/favorites"

type Favorite struct {
	UserID   int
	ThreadID int
}

func ToCore(req Favorite) favorites.Core {
	return favorites.Core{
		UserID:   req.UserID,
		ThreadID: req.ThreadID,
	}
}
