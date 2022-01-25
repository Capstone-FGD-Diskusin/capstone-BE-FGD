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
	"github.com/stretchr/testify/mock"
)

var (
	commentData    b_comments_mock.Data
	commentUsecase comments.Bussiness
	threadUsecase  b_threads_mock.Bussiness

	comment []comments.Core
	thread  []threads.Core
	err1    error
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

	thread = []threads.Core{
		{
			ID:          1,
			Title:       "judul1",
			Description: "Deskripsi1",
			UserID:      1,
			Like:        0,
		},
	}

	err1 = errors.New("tidak dapat menghapus komen orang lain")

	os.Exit(m.Run())
}

func TestAll(t *testing.T) {
	t.Run("valid - add comment", func(t *testing.T) {
		threadUsecase.On("IncrementComment", mock.AnythingOfType("threads.Core")).Return(err1).Once()
		commentData.On("InsertComment", mock.AnythingOfType("comments.Core")).Return(nil).Once()
		err := commentUsecase.AddComment(comment[0])

		// assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, err1)
		// assert.NotEqual(t, err1, nil)
	})

	t.Run("valid - add comment 2", func(t *testing.T) {
		threadUsecase.On("IncrementComment", mock.AnythingOfType("threads.Core")).Return(nil).Once()
		commentData.On("InsertComment", mock.AnythingOfType("comments.Core")).Return(nil).Once()
		err := commentUsecase.AddComment(comment[0])

		// assert.NotEqual(t, len(resp), 0)
		// assert.Equal(t, err, err1)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - get comments thread", func(t *testing.T) {
		commentData.On("SelectCommentsThread", mock.AnythingOfType("comments.Core")).Return(comment, nil).Once()
		resp, err := commentUsecase.GetCommentsThread(comment[0])

		assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - get comments by id", func(t *testing.T) {
		commentData.On("SelectCommentbyId", mock.AnythingOfType("comments.Core")).Return(comment[0], nil)
		resp, err := commentUsecase.GetCommentbyId(comment[0])

		assert.Equal(t, resp, comment[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - delete comment thread", func(t *testing.T) {
		commentData.On("SelectCommentbyId", mock.AnythingOfType("comments.Core")).Return(comment[0], nil)
		cekDataComment, err := commentUsecase.GetCommentbyId(comment[0])
		assert.Equal(t, err, nil)

		threadUsecase.On("GetThreadbyID", mock.AnythingOfType("threads.Core")).Return(thread[0], nil).Once()
		commentData.On("DeleteCommentbyId", mock.AnythingOfType("comments.Core")).Return(nil).Once()

		err = commentUsecase.DeteleCommentThread(comment[0])

		assert.Equal(t, cekDataComment, comment[0])
		// assert.Equal(t, resp, comment[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - delete comment by thread id", func(t *testing.T) {
		commentData.On("DeleteCommentbyThreadId", mock.AnythingOfType("comments.Core")).Return(nil).Once()
		err := commentUsecase.DeleteCommentbyThreadId(comment[0])

		// assert.Equal(t, resp, comment[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - delete comment by thread id 2", func(t *testing.T) {
		commentData.On("DeleteCommentbyThreadId", mock.AnythingOfType("comments.Core")).Return(err1)
		err := commentUsecase.DeleteCommentbyThreadId(comment[0])
		// assert.Equal(t, resp, comment[0])
		assert.Equal(t, err, err1)
	})

	t.Run("valid - get balasan comment id", func(t *testing.T) {
		commentData.On("SelectBalasanCommentbyId", mock.AnythingOfType("comments.Core")).Return(comment, nil).Once()
		resp, err := commentUsecase.GetBalasanCommentbyId(comment[0])
		assert.Equal(t, len(resp), 1)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - get balasan comment id", func(t *testing.T) {
		commentData.On("SelectBalasanCommentbyId", mock.AnythingOfType("comments.Core")).Return(comment, err1)
		resp, err := commentUsecase.GetBalasanCommentbyId(comment[0])
		assert.Equal(t, len(resp), 1)
		assert.Equal(t, err, err1)
	})

	t.Run("valid - search thread", func(t *testing.T) {
		threadUsecase.On("SearchThread", mock.AnythingOfType("threads.Core")).Return(thread, nil)
		commentData.On("SearchThreadbyComment", mock.AnythingOfType("comments.Core")).Return(comment, nil)
		threadUsecase.On("GetThreadbyID", mock.AnythingOfType("threads.Core")).Return(thread[0], nil)
		resp, err := commentUsecase.SearchThread(comment[0])
		assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})
}
