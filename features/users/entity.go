package users

type Core struct {
	ID             int
	Email          string
	Password       string
	Username       string
	Follower       int
	SumLike        int
	SumComment     int
	ProfilePicture string
	Threads        []Thread
}

type Thread struct {
	ID            int
	Title         string
	Description   string
	UserID        int
	Like          int
	JumlahComment int
	ImgUrl        string
}

type Bussiness interface {
	Register(data Core) (err error)
	Login(data Core) (userData Core, token string, isAuth bool, err error)
	GetThreadHome(data Core) (resp Core, err error)
}

type Data interface {
	CreateUser(data Core) (err error)
	CheckEmailPass(email string, pass string) (isAuth bool, user Core, err error)
	SelectDatabyEmail(data Core) (resp Core, err error)
}
