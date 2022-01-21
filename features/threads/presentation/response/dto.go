package response

import "github.com/dragranzer/capstone-BE-FGD/features/threads"

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

func FromCore(res threads.Core) Thread {
	return Thread{
		ID:            res.ID,
		Title:         res.Title,
		Description:   res.Description,
		UserID:        res.UserID,
		Like:          res.Like,
		JumlahComment: res.JumlahComment,
		ImgUrl:        res.ImgUrl,
		CategoryName:  res.CategoryName,
	}
}

func FromCoreSlice(data []threads.Core) []Thread {
	resp := []Thread{}
	for _, value := range data {
		resp = append(resp, FromCore(value))
	}
	return resp
}
