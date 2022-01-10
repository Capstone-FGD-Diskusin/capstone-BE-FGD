package comments

type Core struct {
	ID        int
	Comment   string
	ThreadID  int
	UserID    int
	ImageUrl  string
	CommentID int
}

type Bussiness interface {
	AddComment(data Core) (err error)
}

type Data interface {
	InsertComment(data Core) (err error)
}
