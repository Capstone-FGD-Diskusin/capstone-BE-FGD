package users

type Core struct {
	ID             int
	Email          string
	Password       string
	Username       string
	Alamat         string
	Gender         string
	Phone          string
	Follower       int
	SumLike        int
	SumComment     int
	ProfilePicture string
	Role           string
}

type Bussiness interface {
	Register(data Core) (err error)
	Login(data Core) (userData Core, token string, isAuth bool, err error)
	GetProfileData(data Core) (resp Core, err error)
	IncrementLike(data Core) (err error)
	DecrementLike(data Core) (err error)
	IncrementFol(data Core) (err error)
	DecrementFol(data Core) (err error)
	IncrementFollowing(data Core) (err error)
	DecrementFollowing(data Core) (err error)
	EditDataUser(data Core) (err error)
	DeleteDataUserbyId(data Core) (err error)
}

type Data interface {
	CreateUser(data Core) (err error)
	CheckEmailPass(email string, pass string) (isAuth bool, user Core, err error)
	SelectDatabyEmail(data Core) (resp Core, err error)
	SelectDatabyID(data Core) (resp Core, err error)
	UpdateLikebyOne(data Core) (err error)
	UpdateMinLikebyOne(data Core) (err error)
	UpdateFolbyOne(data Core) (err error)
	UpdateMinFolbyOne(data Core) (err error)
	UpdateFollowingbyOne(data Core) (err error)
	UpdateMinFollowingbyOne(data Core) (err error)
	UpdateDataUser(data Core) (err error)
	DeleteDataUserbyId(data Core) (err error)
}
