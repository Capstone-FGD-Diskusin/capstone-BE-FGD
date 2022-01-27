package data

import "github.com/dragranzer/capstone-BE-FGD/features/likes"

type Like struct {
	UserID   int `gorm:"primary_key"`
	ThreadID int `gorm:"primary_key"`
}

func fromCore(core likes.Core) Like {
	return Like{
		UserID:   core.UserID,
		ThreadID: core.ThreadID,
	}
}
