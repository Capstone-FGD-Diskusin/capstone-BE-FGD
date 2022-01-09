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

func (lu *likesUsecase) UnlikingThread(data likes.Core) (err error) {
	err = lu.likeData.DeleteLike(data)
	if err != nil {
		return err
	}
	thread := threads.Core{
		ID: data.ThreadID,
	}
	err = lu.threadBussiness.DecrementLike(thread)
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
	err = lu.userBussiness.DecrementLike(user)
	return err
}

func (lu *likesUsecase) GetThreadHome(data likes.Core) (resp []likes.Core, err error) {
	temp := threads.Core{
		OwnerID: data.UserID,
		Page:    data.Page,
	}
	threads, err := lu.threadBussiness.GetThreadHome(temp)
	for _, value := range threads {
		check := likes.Core{
			UserID:   data.UserID,
			ThreadID: value.ID,
		}
		isLiked, err := lu.likeData.CheckLiked(check)
		if err != nil {
			return resp, err
		}
		thread := likes.Thread{
			ID:            value.ID,
			Title:         value.Title,
			Description:   value.Description,
			UserID:        value.UserID,
			Like:          value.Like,
			JumlahComment: value.JumlahComment,
			ImgUrl:        value.ImgUrl,
			IsLiked:       isLiked,
		}
		resp = append(resp, likes.Core{
			Thread: thread,
		})
	}
	return
}
