package followers

type Core struct {
	FollowingID  int
	FollowedID   int
	NameFollowed string
}

type Bussiness interface {
	Follow(data Core) (err error)
	Unfollow(data Core) (err error)
	GetFollowing(data Core) (resp []Core, err error)
}

type Data interface {
	InsertFollow(data Core) (err error)
	DeleteFollow(data Core) (err error)
	SelectFollowing(data Core) (resp []Core, err error)
}
