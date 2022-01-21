package request

import "github.com/dragranzer/capstone-BE-FGD/features/messages"

type Message struct {
	ID           int
	Text         string `json:"text" form:"text"`
	ThreadID     int    `json:"thread_id" form:"thread_id"`
	CategoryName string `json:"category_name" form:"category_name"`
	CommentID    int    `json:"comment_id" form:"comment_id"`
	ModeratorID  int
	AdminID      int `json:"admin_id" form:"admin_id"`
}

func ToCore(req Message) messages.Core {
	return messages.Core{
		Text:         req.Text,
		ThreadID:     req.ThreadID,
		CategoryName: req.CategoryName,
		CommentID:    req.CommentID,
		ModeratorID:  req.ModeratorID,
		AdminID:      req.AdminID,
	}
}
