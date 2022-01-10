package bussiness

import "github.com/dragranzer/capstone-BE-FGD/features/comments"

type commentsUsecase struct {
	commentData comments.Data
}

func NewCommentBussiness(cD comments.Data) comments.Bussiness {
	return &commentsUsecase{
		commentData: cD,
	}
}

func (cU *commentsUsecase) AddComment(data comments.Core) (err error) {
	err = cU.commentData.InsertComment(data)
	return
}
