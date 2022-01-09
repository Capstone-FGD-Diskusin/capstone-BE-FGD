package bussiness

import (
	"github.com/dragranzer/capstone-BE-FGD/features/likes"
	"github.com/dragranzer/capstone-BE-FGD/features/threads"
	"github.com/dragranzer/capstone-BE-FGD/features/users"
)

type likesUsecase struct {
	likeData        likes.Data
	userBussiness   users.Bussiness
	threadBussiness threads.Bussiness
}

func NewLikeBussiness(lD likes.Data, uB users.Bussiness, tB threads.Bussiness) likes.Bussiness {
	return &likesUsecase{
		likeData:        lD,
		userBussiness:   uB,
		threadBussiness: tB,
	}
}

func (lu *likesUsecase) LikingThread(data likes.Core) (err error) {
	err = lu.likeData.InsertLike(data)
	if err != nil {
		return err
	}
	thread := threads.Core{
		ID: data.ThreadID,
	}
	err = lu.threadBussiness.IncrementLike(thread)
	if err != nil {
		return err
	}
	thread, err = lu.threadBussiness.GetThreadbyID(thread)
	if err != nil {
		return err
	}
	user := users.Core{
		ID: thread.UserID,
	}
	err = lu.userBussiness.IncrementLike(user)
	return err
}
