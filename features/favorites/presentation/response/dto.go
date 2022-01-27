package response

import "github.com/dragranzer/capstone-BE-FGD/features/favorites"

type Favorite struct {
	UserID        int
	ThreadID      int
	ID            int
	Title         string
	Description   string
	Like          int
	JumlahComment int
	ImgUrl        string
	CategoryName  string
	UserName      string
}

func FromCore(res favorites.Core) Favorite {
	return Favorite{
		UserID:        res.UserID,
		ThreadID:      res.ThreadID,
		ID:            res.Thread.ID,
		Title:         res.Thread.Title,
		Description:   res.Thread.Description,
		Like:          res.Thread.Like,
		JumlahComment: res.Thread.JumlahComment,
		ImgUrl:        res.Thread.ImgUrl,
		CategoryName:  res.Thread.CategoryName,
		UserName:      res.Thread.UserName,
	}
}

func FromCoreSlice(data []favorites.Core) []Favorite {
	resp := []Favorite{}
	for _, value := range data {
		resp = append(resp, FromCore(value))
	}
	return resp
}
