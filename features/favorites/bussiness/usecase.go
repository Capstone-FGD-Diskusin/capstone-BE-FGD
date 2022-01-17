package bussiness

import (
	"github.com/dragranzer/capstone-BE-FGD/features/comments"
	"github.com/dragranzer/capstone-BE-FGD/features/favorites"
	"github.com/dragranzer/capstone-BE-FGD/features/threads"
	"github.com/dragranzer/capstone-BE-FGD/features/users"
)

type favoritesUsecase struct {
	threadBussiness  threads.Bussiness
	UserBussiness    users.Bussiness
	commentBussiness comments.Bussiness
	favoriteData     favorites.Data
}

func NewFavoriteBussiness(tB threads.Bussiness, uB users.Bussiness, cB comments.Bussiness, fD favorites.Data) favorites.Bussiness {
	return &favoritesUsecase{
		threadBussiness:  tB,
		UserBussiness:    uB,
		commentBussiness: cB,
		favoriteData:     fD,
	}
}

func (fu *favoritesUsecase) DeleteThreadbyId(data favorites.Core) (err error) {
	comment_core := comments.Core{
		ThreadID: data.ThreadID,
	}
	err = fu.commentBussiness.DeleteCommentbyThreadId(comment_core)
	if err != nil {
		return err
	}
	thread_core := threads.Core{
		ID: data.ThreadID,
	}
	err = fu.threadBussiness.DeleteThreadbyId(thread_core)
	return
}
