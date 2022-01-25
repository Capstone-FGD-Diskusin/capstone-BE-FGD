package bussiness_test

import (
	"errors"
	"os"
	"testing"

	"github.com/dragranzer/capstone-BE-FGD/features/comments"
	b_comments "github.com/dragranzer/capstone-BE-FGD/features/comments/bussiness"
	b_comments_mock "github.com/dragranzer/capstone-BE-FGD/features/comments/mocks"
	"github.com/dragranzer/capstone-BE-FGD/features/threads"
	b_threads_mock "github.com/dragranzer/capstone-BE-FGD/features/threads/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	commentData    b_comments_mock.Data
	commentUsecase comments.Bussiness
	threadUsecase  b_threads_mock.Bussiness

	comment  []comments.Core
	comment2 []comments.Core
	thread   []threads.Core
	thread2  []threads.Core
	err1     error
)

func TestMain(m *testing.M) {
	commentUsecase = b_comments.NewCommentBussiness(&commentData, &threadUsecase)

	comment = []comments.Core{
		{
			ID:       1,
			Comment:  "Keren banget",
			ThreadID: 1,
			UserID:   1,
			ImageUrl: "www.image.com",
		},
	}

	comment2 = []comments.Core{
		{
			ID: 0,
			// Comment:  "Keren banget",
			// ThreadID: 1,
			// UserID:   1,
			// ImageUrl: "www.image.com",
		},
	}

	thread = []threads.Core{
		{
			ID: 1,
			// Title:       "judul1",
			// Description: "Deskripsi1",
			// UserID:      1,
			// Like:        0,
		},
	}

	thread2 = []threads.Core{
		{
			ID: 0,
		},
	}

	err1 = errors.New("tidak dapat menghapus komen orang lain")

	os.Exit(m.Run())
}

func TestAll(t *testing.T) {
	t.Run("valid - add comment", func(t *testing.T) {
		threadUsecase.On("IncrementComment", thread[0]).Return(err1).Once()
		commentData.On("InsertComment", comment[0]).Return(nil).Once()
		err := commentUsecase.AddComment(comment[0])

		// assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, err1)
		// assert.NotEqual(t, err1, nil)
	})

	t.Run("valid - add comment 2", func(t *testing.T) {
		threadUsecase.On("IncrementComment", thread[0]).Return(nil).Once()
		commentData.On("InsertComment", comment[0]).Return(nil).Once()
		err := commentUsecase.AddComment(comment[0])

		// assert.NotEqual(t, len(resp), 0)
		// assert.Equal(t, err, err1)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - get comments thread", func(t *testing.T) {
		commentData.On("SelectCommentsThread", comment[0]).Return(comment, nil).Once()
		resp, err := commentUsecase.GetCommentsThread(comment[0])

		assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - get comments by id", func(t *testing.T) {
		commentData.On("SelectCommentbyId", comment[0]).Return(comment[0], nil)
		resp, err := commentUsecase.GetCommentbyId(comment[0])

		assert.Equal(t, resp, comment[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - delete comment thread", func(t *testing.T) {
		commentData.On("SelectCommentbyId", comment[0]).Return(comment[0], nil)
		cekDataComment, err := commentUsecase.GetCommentbyId(comment[0])
		assert.Equal(t, err, nil)

		threadUsecase.On("GetThreadbyID", thread[0]).Return(thread[0], nil).Once()
		commentData.On("DeleteCommentbyId", comment[0]).Return(nil).Once()

		err = commentUsecase.DeteleCommentThread(comment[0])

		assert.Equal(t, cekDataComment, comment[0])
		// assert.Equal(t, resp, comment[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - delete comment by thread id", func(t *testing.T) {
		commentData.On("DeleteCommentbyThreadId", comment[0]).Return(nil).Once()
		err := commentUsecase.DeleteCommentbyThreadId(comment[0])

		// assert.Equal(t, resp, comment[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - delete comment by thread id 2", func(t *testing.T) {
		commentData.On("DeleteCommentbyThreadId", comment[0]).Return(err1)
		err := commentUsecase.DeleteCommentbyThreadId(comment[0])
		// assert.Equal(t, resp, comment[0])
		assert.Equal(t, err, err1)
	})

	t.Run("valid - get balasan comment id", func(t *testing.T) {
		commentData.On("SelectBalasanCommentbyId", comment[0]).Return(comment, nil).Once()
		resp, err := commentUsecase.GetBalasanCommentbyId(comment[0])
		assert.Equal(t, len(resp), 1)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - get balasan comment id", func(t *testing.T) {
		commentData.On("SelectBalasanCommentbyId", comment[0]).Return(comment, err1)
		resp, err := commentUsecase.GetBalasanCommentbyId(comment[0])
		assert.Equal(t, len(resp), 1)
		assert.Equal(t, err, err1)
	})

	t.Run("valid - search thread", func(t *testing.T) {
		threadUsecase.On("SearchThread", thread2[0]).Return(thread2, nil)
		commentData.On("SearchThreadbyComment", comment2[0]).Return(comment2, nil)
		threadUsecase.On("GetThreadbyID", thread[0]).Return(thread[0], nil)
		resp, err := commentUsecase.SearchThread(comment[0])
		assert.Equal(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})
}
