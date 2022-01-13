package data

import "github.com/dragranzer/capstone-BE-FGD/features/comments"

type Comment struct {
	ID        int
	Comment   string
	UserID    int
	ThreadID  int
	ImageUrl  string
	CommentID int
}

func FromCore(data comments.Core) Comment {
	return Comment{
		Comment:   data.Comment,
		ThreadID:  data.ThreadID,
		UserID:    data.UserID,
		ImageUrl:  data.ImageUrl,
		CommentID: data.CommentID,
	}
}
