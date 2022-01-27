package data

import "github.com/dragranzer/capstone-BE-FGD/features/favorites"

type Favorite struct {
	UserID   int `gorm:"primary_key"`
	ThreadID int `gorm:"primary_key"`
}

func FromCore(data favorites.Core) Favorite {
	return Favorite{
		ThreadID: data.ThreadID,
		UserID:   data.UserID,
	}
}

func (a *Favorite) toCore() favorites.Core {
	return favorites.Core{
		ThreadID: a.ThreadID,
		UserID:   a.UserID,
	}
}

func ToCoreSlice(data []Favorite) []favorites.Core {
	resp := []favorites.Core{}
	for _, value := range data {
		resp = append(resp, value.toCore())
	}
	return resp
}
