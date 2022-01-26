package favorites

type Core struct {
	UserID   int
	ThreadID int
}

type Bussiness interface {
	DeleteThreadbyId(data Core) (err error)
	InsertFavorite(data Core) (err error)
	DeleteFavorite(data Core) (err error)
}

type Data interface {
	DeleteFavoritebyThreadId(data Core) (err error)
	AddFavorite(data Core) (err error)
	DeleteFavorite(data Core) (err error)
}
