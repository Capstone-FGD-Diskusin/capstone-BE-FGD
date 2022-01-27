package favorites

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
	Like          int
	JumlahComment int
	ImgUrl        string
	CategoryName  string
	UserName      string
}

type Bussiness interface {
	DeleteThreadbyId(data Core) (err error)
	InsertFavorite(data Core) (err error)
	DeleteFavorite(data Core) (err error)
	GetAllFavorite(data Core) (resp []Core, err error)
}

type Data interface {
	DeleteFavoritebyThreadId(data Core) (err error)
	AddFavorite(data Core) (err error)
	DeleteFavorite(data Core) (err error)
	SelectAllFavorite(data Core) (resp []Core, err error)
}
