package bussiness

import (
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
	return
}

func (cU *commentsUsecase) DeteleCommentThread(data comments.Core) (err error) {
	err = cU.commentData.DeleteCommentbyId(data)
	return
}
