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

type Thread struct {
	ID            int
	Title         string
	Description   string
	UserID        int
	Like          int
	JumlahComment int
	ImgUrl        string
	CategoryName  string
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

func FromCoreThread(res comments.Core) Thread {
	return Thread{
		ID:           res.Thread.ID,
		Title:        res.Thread.Title,
		Description:  res.Thread.Description,
		UserID:       res.Thread.UserID,
		ImgUrl:       res.Thread.ImgUrl,
		CategoryName: res.Thread.CategoryName,
	}
}

func FromCoreSliceThread(data []comments.Core) []Thread {
	resp := []Thread{}
	for _, value := range data {
		resp = append(resp, FromCoreThread(value))
	}
	return resp
}
