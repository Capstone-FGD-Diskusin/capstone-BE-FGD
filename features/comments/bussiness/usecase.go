package bussiness

import (
	"errors"
	"fmt"

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

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
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
	cekDataComment, err := cU.GetCommentbyId(data)
	if err != nil {
		return err
	}

	threadCore := threads.Core{
		ID: cekDataComment.ThreadID,
	}
	cekDataThread, err := cU.threadBussiness.GetThreadbyID(threadCore)
	if err != nil {
		return err
	}

	if cekDataComment.UserID != data.UserID {
		if cekDataThread.UserID != data.UserID {
			err = errors.New("tidak dapat menghapus komen orang lain")
			return err
		}
	}
	err = cU.commentData.DeleteCommentbyId(data)
	return err
}

func (cU *commentsUsecase) DeleteCommentbyThreadId(data comments.Core) (err error) {
	err = cU.commentData.DeleteCommentbyThreadId(data)
	if err != nil {
		return err
	}
	return err
}

func (cU *commentsUsecase) GetBalasanCommentbyId(data comments.Core) (resp []comments.Core, err error) {
	resp, err = cU.commentData.SelectBalasanCommentbyId(data)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func (cU *commentsUsecase) SearchThread(data comments.Core) (resp []comments.Core, err error) {
	threadCore := threads.Core{
		Search: data.Search,
	}
	listThread, err := cU.threadBussiness.SearchThread(threadCore)

	if err != nil {
		return resp, err
	}

	listThreadID := []int{}
	// fmt.Println(listThread)
	for _, value := range listThread {
		listThreadID = append(listThreadID, value.ID)
	}

	commentCore := comments.Core{
		Search: data.Search,
	}
	listThread2, err := cU.commentData.SearchThreadbyComment(commentCore)
	for _, value := range listThread2 {
		listThreadID = append(listThreadID, value.ThreadID)
	}

	listThreadID = unique(listThreadID)
	fmt.Println(listThreadID)

	for _, value := range listThreadID {
		if value == 0 {
			continue
		}

		threadCore := threads.Core{
			ID: value,
		}

		threadCore, err = cU.threadBussiness.GetThreadbyID(threadCore)
		thread := comments.Thread{
			ID:           threadCore.ID,
			Title:        threadCore.Title,
			Description:  threadCore.Description,
			UserID:       threadCore.UserID,
			ImgUrl:       threadCore.ImgUrl,
			CategoryName: threadCore.CategoryName,
		}
		resp = append(resp, comments.Core{
			Thread: thread,
		})
	}

	if err != nil {
		return resp, err
	}
	return resp, err
}
