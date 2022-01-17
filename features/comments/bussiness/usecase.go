package bussiness

import (
	"errors"

	"github.com/dragranzer/capstone-BE-FGD/features/comments"
	"github.com/dragranzer/capstone-BE-FGD/features/threads"
)

type commentsUsecase struct {
	commentData     comments.Data
	threadBussiness threads.Bussiness
}

func NewCommentBussiness(cD comments.Data, tB threads.Bussiness) comments.Bussiness {
	return &commentsUsecase{
		commentData:     cD,
		threadBussiness: tB,
	}
}

func (cU *commentsUsecase) AddComment(data comments.Core) (err error) {
	thread := threads.Core{
		ID: data.ThreadID,
	}
	err = cU.threadBussiness.IncrementComment(thread)
	if err != nil {
		return err
	}
	err = cU.commentData.InsertComment(data)
	return
}

func (cU *commentsUsecase) GetCommentsThread(data comments.Core) (resp []comments.Core, err error) {
	resp, err = cU.commentData.SelectCommentsThread(data)
	if err != nil {
		return resp, err
	}
	return resp, err
}
func (cU *commentsUsecase) GetCommentbyId(data comments.Core) (resp comments.Core, err error) {
	resp, err = cU.commentData.SelectCommentbyId(data)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func (cU *commentsUsecase) DeteleCommentThread(data comments.Core) (err error) {
	cekData, err := cU.GetCommentbyId(data)
	if err != nil {
		return err
	}
	if cekData.UserID != data.UserID {
		err = errors.New("tidak dapat menghapus komen orang lain")
		return err
	}
	err = cU.commentData.DeleteCommentbyId(data)
	return err
}
