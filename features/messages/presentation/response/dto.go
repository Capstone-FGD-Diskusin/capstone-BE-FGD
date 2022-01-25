package response

import "github.com/dragranzer/capstone-BE-FGD/features/messages"

type Message struct {
	ID           int
	Text         string
	ThreadID     int
	ThreadTitle  string
	CategoryName string
	CommentID    int
	Comment      string
	AdminID      int
	AdminName    string
}

func FromCore(res messages.Core) Message {
	return Message{
		ID:           res.ID,
		Text:         res.Text,
		ThreadID:     res.ThreadID,
		ThreadTitle:  res.ThreadTitle,
		CategoryName: res.CategoryName,
		CommentID:    res.CommentID,
		Comment:      res.Comment,
		AdminID:      res.AdminID,
		AdminName:    res.AdminName,
	}
}

func FromCoreSlice(data []messages.Core) []Message {
	resp := []Message{}
	for _, value := range data {
		resp = append(resp, FromCore(value))
	}
	return resp
}
