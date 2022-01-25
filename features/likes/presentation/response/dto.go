package response

import "github.com/dragranzer/capstone-BE-FGD/features/likes"

type Thread struct {
	ID            int
	Title         string
	Description   string
	UserID        int
	Like          int
	JumlahComment int
	ImgUrl        string
	IsLiked       bool
	CategoryName  string
	UserName      string
}

func FromCore(res likes.Core) Thread {
	return Thread{
		ID:            res.Thread.ID,
		Title:         res.Thread.Title,
		Description:   res.Thread.Description,
		UserID:        res.Thread.UserID,
		Like:          res.Thread.Like,
		JumlahComment: res.Thread.JumlahComment,
		ImgUrl:        res.Thread.ImgUrl,
		IsLiked:       res.Thread.IsLiked,
		CategoryName:  res.Thread.CategoryName,
		UserName:      res.Thread.UserName,
	}
}

func FromCoreSlice(data []likes.Core) []Thread {
	resp := []Thread{}
	for _, value := range data {
		resp = append(resp, FromCore(value))
	}
	return resp
}
