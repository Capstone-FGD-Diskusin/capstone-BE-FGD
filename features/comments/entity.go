package comments

type Core struct {
	ID        int
	Comment   string
	ThreadID  int
	UserID    int
	ImageUrl  string
	CommentID int
	Page      int
	Search    string
	Thread    Thread
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
	AddComment(data Core) (err error)
	GetCommentsThread(data Core) (resp []Core, err error)
	DeteleCommentThread(data Core) (err error)
	GetCommentbyId(data Core) (resp Core, err error)
	DeleteCommentbyThreadId(data Core) (err error)
	GetBalasanCommentbyId(data Core) (resp []Core, err error)
	SearchThread(data Core) (resp []Core, err error)
}

type Data interface {
	InsertComment(data Core) (err error)
	SelectCommentsThread(data Core) (resp []Core, err error)
	DeleteCommentbyId(data Core) (err error)
	SelectCommentbyId(data Core) (resp Core, err error)
	DeleteCommentbyThreadId(data Core) (err error)
	SelectBalasanCommentbyId(data Core) (resp []Core, err error)
}
