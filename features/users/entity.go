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
}

type Data interface {
	CreateUser(data Core) (err error)
}
