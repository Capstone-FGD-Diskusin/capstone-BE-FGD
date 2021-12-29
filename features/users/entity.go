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
}

type Bussiness interface {
	Register(data Core) (err error)
	Login(email string, pass string) (userData Core, token string, isAuth bool, err error)
}

type Data interface {
	CreateUser(data Core) (err error)
	CheckEmailPass(email string, pass string) (isAuth bool, user Core, err error)
	SelectDatabyEmail(email string) (resp Core, err error)
}
