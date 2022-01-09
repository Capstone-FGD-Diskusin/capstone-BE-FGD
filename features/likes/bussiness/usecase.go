package bussiness

import (
	"github.com/dragranzer/capstone-BE-FGD/features/likes"
	"github.com/dragranzer/capstone-BE-FGD/features/threads"
	"github.com/dragranzer/capstone-BE-FGD/features/users"
)

type likesUsecase struct {
	likeData        likes.Data
	userBussiness   users.Bussiness
	ThreadBussiness threads.Bussiness
}

func NewLikeBussiness(lD likes.Data, uB users.Bussiness, tB threads.Bussiness) likes.Bussiness {
	return &likesUsecase{
		likeData:        lD,
		userBussiness:   uB,
		ThreadBussiness: tB,
	}
}

func (lu *likesUsecase) LikingThread(data likes.Core) (err error) {
	err = lu.likeData.InsertLike(data)
	return err
}
