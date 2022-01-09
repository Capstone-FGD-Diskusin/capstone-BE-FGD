package data

import "github.com/dragranzer/capstone-BE-FGD/features/likes"

type Like struct {
	UserID   int
	ThreadID int
}

func fromCore(core likes.Core) Like {
	return Like{
		UserID:   core.UserID,
		ThreadID: core.ThreadID,
	}
}
