package likes

type Core struct {
	UserID   int
	ThreadID int
	Page     int
	Thread   Thread
}

type Thread struct {
	ID            int
	Title         string
	Description   string
	UserID        int
	Like          int
	JumlahComment int
	ImgUrl        string
	IsLiked       bool
}

type Bussiness interface {
	LikingThread(data Core) (err error)
	UnlikingThread(data Core) (err error)
	GetThreadHome(data Core) (resp []Core, err error)
}

type Data interface {
	InsertLike(data Core) (err error)
	DeleteLike(data Core) (err error)
	CheckLiked(data Core) (isLiked bool, err error)
}
