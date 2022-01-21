package data

import "github.com/dragranzer/capstone-BE-FGD/features/messages"

type Message struct {
	ID          int
	Text        string
	ThreadID    int
	CategoryID  int
	CommentID   int
	ModeratorID int
	AdminID     int
}

func fromCore(core messages.Core) Message {
	return Message{
		ID:          core.ID,
		Text:        core.Text,
		ThreadID:    core.ThreadID,
		CategoryID:  core.CategoryID,
		CommentID:   core.CommentID,
		ModeratorID: core.ModeratorID,
		AdminID:     core.AdminID,
	}
}

func (a *Message) toCore() messages.Core {
	return messages.Core{
		ID:          int(a.ID),
		Text:        a.Text,
		ThreadID:    a.ThreadID,
		CategoryID:  a.CategoryID,
		CommentID:   a.CommentID,
		ModeratorID: a.ModeratorID,
		AdminID:     a.AdminID,
	}
}

func ToCoreSlice(data []Message) []messages.Core {
	resp := []messages.Core{}
	for _, value := range data {
		resp = append(resp, value.toCore())
	}
	return resp
}
