package request

import "github.com/dragranzer/capstone-BE-FGD/features/comments"

type Comment struct {
	ID        int    `json:"id" form:"id"`
	Comment   string `json:"comment" form:"comment"`
	UserID    int
	ThreadID  int    `json:"thread_id" form:"thread_id"`
	ImageUrl  string `json:"image_url" form:"image_url"`
	CommentID int    `json:"comment_id" form:"comment_id"`
	Page      int    `json:"page" form:"page"`
	Search    string `json:"search" form:"search"`
}

func ToCore(req Comment) comments.Core {
	return comments.Core{
		ID:        req.ID,
		Comment:   req.Comment,
		UserID:    req.UserID,
		ThreadID:  req.ThreadID,
		ImageUrl:  req.ImageUrl,
		CommentID: req.CommentID,
	}
}
