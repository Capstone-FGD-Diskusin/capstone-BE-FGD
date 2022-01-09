package likes

type Core struct {
	UserID   int
	ThreadID int
}

type Bussiness interface {
	LikingThread(data Core) (err error)
	UnlikingThread(data Core) (err error)
}

type Data interface {
	InsertLike(data Core) (err error)
	DeleteLike(data Core) (err error)
}
