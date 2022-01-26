package data

import "github.com/dragranzer/capstone-BE-FGD/features/favorites"

type Favorite struct {
	UserID   int
	ThreadID int
}

func FromCore(data favorites.Core) Favorite {
	return Favorite{
		ThreadID: data.ThreadID,
		UserID:   data.UserID,
	}
}
