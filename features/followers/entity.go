package followers

type Core struct {
	FollowingID int
	FollowedID  int
}

type Bussiness interface {
	Follow(data Core) (err error)
}

type Data interface {
	InsertFollow(data Core) (err error)
}
