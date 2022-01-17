package favorites

type Core struct {
	UserID   int
	ThreadID int
}

type Bussiness interface {
	DeleteThreadbyId(data Core) (err error)
}

type Data interface {
}
