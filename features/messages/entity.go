package messages

type Core struct {
	ID            int
	Text          string
	ThreadID      int
	ThreadTitle   string
	CategoryID    int
	CategoryName  string
	CommentID     int
	Comment       string
	ModeratorID   int
	ModeratorName string
	AdminID       int
	AdminName     string
}

type Bussiness interface {
	ReportToAdmin(data Core) (err error)
	GetMessagesbyAdminID(data Core) (resp []Core, err error)
	DeleteMessagesbyId(data Core) (err error)
}

type Data interface {
	InsertMessages(data Core) (err error)
	SelectMessagesbyAdminID(data Core) (resp []Core, err error)
	DeleteMessagesbyId(data Core) (err error)
}
