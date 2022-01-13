package request

import "github.com/dragranzer/capstone-BE-FGD/features/comments"

type Comment struct {
	Comment   string `json:"comment" form:"comment"`
	UserID    int
	ThreadID  int    `json:"thread_id" form:"thread_id"`
	ImageUrl  string `json:"image_url" form:"image_url"`
	CommentID int    `json:"comment_id" form:"comment_id"`
	Page      int    `json:"page" form:"page"`
}

func ToCore(req Comment) comments.Core {
	return comments.Core{
		Comment:   req.Comment,
		UserID:    req.UserID,
		ThreadID:  req.ThreadID,
		ImageUrl:  req.ImageUrl,
		CommentID: req.CommentID,
	}
}
