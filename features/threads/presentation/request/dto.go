package request

import "github.com/dragranzer/capstone-BE-FGD/features/threads"

type Thread struct {
	Title        string `json:"title" form:"title"`
	Description  string `json:"description" form:"description"`
	UserID       int
	ImgUrl       string `json:"img_url" form:"img_url"`
	CategoryName string `json:"category_name" form:"category_name"`
}

type Request struct {
	UserID int `json:"user_id" form:"user_id"`
	Page   int `json:"page" form:"page"`
}

func ToCore(req Request) threads.Core {
	return threads.Core{
		OwnerID: req.UserID,
		Page:    req.Page,
		UserID:  req.UserID,
	}
}

func ToCoreThread(req Thread) threads.Core {
	return threads.Core{
		Title:        req.Title,
		Description:  req.Description,
		UserID:       req.UserID,
		ImgUrl:       req.ImgUrl,
		CategoryName: req.CategoryName,
	}
}
