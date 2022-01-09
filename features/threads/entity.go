package threads

type Core struct {
	ID             int
	Title          string
	Description    string
	UserID         int
	Like           int
	JumlahComment  int
	ImgUrl         string
	ListFollowedID []int
	OwnerID        int
	Page           int
}

type Bussiness interface {
	GetThreadHome(data Core) (resp []Core, err error)
	AddThread(data Core) (err error)
	GetThreadbyID(data Core) (resp Core, err error)
	IncrementLike(data Core) (err error)
}

type Data interface {
	SelectThreadHome(data Core) (resp []Core, err error)
	InsertThread(data Core) (err error)
	SelectThreadbyID(data Core) (resp Core, err error)
	UpdateLikebyOne(data Core) (err error)
}
