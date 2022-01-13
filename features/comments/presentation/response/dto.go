package response

import "github.com/dragranzer/capstone-BE-FGD/features/comments"

type Comment struct {
	ID        int
	Comment   string
	ThreadID  int
	UserID    int
	ImageUrl  string
	CommentID int
}

func FromCore(res comments.Core) Comment {
	return Comment{
		ID:        res.ID,
		Comment:   res.Comment,
		ThreadID:  res.ThreadID,
		UserID:    res.UserID,
		CommentID: res.CommentID,
		ImageUrl:  res.ImageUrl,
	}
}

func FromCoreSlice(data []comments.Core) []Comment {
	resp := []Comment{}
	for _, value := range data {
		resp = append(resp, FromCore(value))
	}
	return resp
}
