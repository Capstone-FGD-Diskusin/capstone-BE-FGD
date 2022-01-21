package data

import "github.com/dragranzer/capstone-BE-FGD/features/messages"

type Message struct {
	ID            int
	Text          string
	ThreadID      int
	CategoryName  string
	CommentID     int
	ModeratorName string
	AdminID       int
}

func fromCore(core messages.Core) Message {
	return Message{
		ID:            core.ID,
		Text:          core.Text,
		ThreadID:      core.ThreadID,
		CategoryName:  core.CategoryName,
		CommentID:     core.CommentID,
		ModeratorName: core.ModeratorName,
		AdminID:       core.AdminID,
	}
}

func (a *Message) toCore() messages.Core {
	return messages.Core{
		ID:            int(a.ID),
		Text:          a.Text,
		ThreadID:      a.ThreadID,
		CategoryName:  a.CategoryName,
		CommentID:     a.CommentID,
		ModeratorName: a.ModeratorName,
		AdminID:       a.AdminID,
	}
}

func ToCoreSlice(data []Message) []messages.Core {
	resp := []messages.Core{}
	for _, value := range data {
		resp = append(resp, value.toCore())
	}
	return resp
}
