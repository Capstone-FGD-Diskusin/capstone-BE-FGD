package data

import (
	"time"

	"github.com/dragranzer/capstone-BE-FGD/features/threads"
)

type Thread struct {
	ID            int
	Title         string
	Description   string
	UserID        int
	Like          int
	JumlahComment int
	ImgUrl        string
	CategoryID    int
	CreatedAt     time.Time
}

func FromCore(data threads.Core) Thread {
	return Thread{
		Title:       data.Title,
		Description: data.Description,
		UserID:      data.UserID,
		ImgUrl:      data.ImgUrl,
		CategoryID:  data.CategoryID,
	}
}

func ToCore(thread Thread) threads.Core {
	return threads.Core{
		ID:            thread.ID,
		Title:         thread.Title,
		Description:   thread.Description,
		UserID:        thread.UserID,
		Like:          thread.Like,
		JumlahComment: thread.JumlahComment,
		ImgUrl:        thread.ImgUrl,
		CategoryID:    thread.CategoryID,
	}
}

func ToCoreSlice(data []Thread) []threads.Core {
	resp := []threads.Core{}
	for _, value := range data {
		resp = append(resp, ToCore(value))
	}
	return resp
}
