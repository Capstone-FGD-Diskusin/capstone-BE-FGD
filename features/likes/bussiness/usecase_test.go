package bussiness_test

import (
	"errors"
	"os"
	"testing"

	"github.com/dragranzer/capstone-BE-FGD/features/likes"
	"github.com/dragranzer/capstone-BE-FGD/features/threads"
	b_threads_mock "github.com/dragranzer/capstone-BE-FGD/features/threads/mocks"
	"github.com/dragranzer/capstone-BE-FGD/features/users"
	b_users_mock "github.com/dragranzer/capstone-BE-FGD/features/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	b_likes "github.com/dragranzer/capstone-BE-FGD/features/likes/bussiness"
	b_likes_mock "github.com/dragranzer/capstone-BE-FGD/features/likes/mocks"
)

var (
	userUsecase   b_users_mock.Bussiness
	threadUsecase b_threads_mock.Bussiness
	likeData      b_likes_mock.Data
	likeUsecase   likes.Bussiness

	like   []likes.Core
	thread []threads.Core
	user   []users.Core

	err1 error
)

func TestMain(m *testing.M) {
	likeUsecase = b_likes.NewLikeBussiness(&likeData, &userUsecase, &threadUsecase)

	like = []likes.Core{
		{
			UserID:   1,
			ThreadID: 1,
		},
	}

	thread = []threads.Core{
		{
			ID:          1,
			Title:       "judul 1",
			Description: "desc1",
			UserID:      1,
		},
	}

	user = []users.Core{
		{
			ID:       1,
			Email:    "email@email.com",
			Username: "saya",
		},
	}

	err1 = errors.New("tidak dapat menghapus komen orang lain")

	os.Exit(m.Run())
}

func TestAll(t *testing.T) {
	t.Run("valid - liking thread", func(t *testing.T) {
		threadUsecase.On("IncrementLike", mock.AnythingOfType("threads.Core")).Return(nil).Once()
		threadUsecase.On("GetThreadbyID", mock.AnythingOfType("threads.Core")).Return(thread[0], nil).Once()
		userUsecase.On("IncrementLike", mock.AnythingOfType("users.Core")).Return(nil).Once()
		likeData.On("InsertLike", mock.AnythingOfType("likes.Core")).Return(nil).Once()
		err := likeUsecase.LikingThread(like[0])

		// assert.Equal(t, resp, comment[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - liking thread", func(t *testing.T) {
		threadUsecase.On("IncrementLike", mock.AnythingOfType("threads.Core")).Return(err1).Once()
		threadUsecase.On("GetThreadbyID", mock.AnythingOfType("threads.Core")).Return(thread[0], nil).Once()
		userUsecase.On("IncrementLike", mock.AnythingOfType("users.Core")).Return(nil).Once()
		likeData.On("InsertLike", mock.AnythingOfType("likes.Core")).Return(nil).Once()
		err := likeUsecase.LikingThread(like[0])

		// assert.Equal(t, resp, comment[0])
		assert.Equal(t, err, err1)
	})

	t.Run("valid - unliking thread", func(t *testing.T) {
		likeData.On("DeleteLike", mock.AnythingOfType("likes.Core")).Return(nil).Once()
		threadUsecase.On("DecrementLike", mock.AnythingOfType("threads.Core")).Return(nil).Once()
		threadUsecase.On("GetThreadbyID", mock.AnythingOfType("threads.Core")).Return(thread[0], nil).Once()
		userUsecase.On("DecrementLike", mock.AnythingOfType("users.Core")).Return(nil).Once()

		err := likeUsecase.UnlikingThread(like[0])

		// assert.Equal(t, resp, comment[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - unliking thread", func(t *testing.T) {
		likeData.On("DeleteLike", mock.AnythingOfType("likes.Core")).Return(err1).Once()
		threadUsecase.On("DecrementLike", mock.AnythingOfType("threads.Core")).Return(nil).Once()
		threadUsecase.On("GetThreadbyID", mock.AnythingOfType("threads.Core")).Return(thread[0], nil).Once()
		userUsecase.On("DecrementLike", mock.AnythingOfType("users.Core")).Return(nil).Once()

		err := likeUsecase.UnlikingThread(like[0])

		// assert.Equal(t, resp, comment[0])
		assert.Equal(t, err, err1)
	})

	t.Run("valid - get thread home", func(t *testing.T) {

		threadUsecase.On("GetThreadHome", mock.AnythingOfType("threads.Core")).Return(thread, nil).Once()
		likeData.On("CheckLiked", mock.AnythingOfType("likes.Core")).Return(true, nil).Once()

		resp, err := likeUsecase.GetThreadHome(like[0])

		assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - get thread home", func(t *testing.T) {

		threadUsecase.On("GetThreadHome", mock.AnythingOfType("threads.Core")).Return(thread, nil).Once()
		likeData.On("CheckLiked", mock.AnythingOfType("likes.Core")).Return(true, err1).Once()

		resp, err := likeUsecase.GetThreadHome(like[0])

		assert.Equal(t, len(resp), 0)
		assert.Equal(t, err, err1)
	})

	t.Run("valid - get thread home", func(t *testing.T) {
		likeData.On("DeleteLikebyThreadId", mock.AnythingOfType("likes.Core")).Return(nil).Once()

		err := likeUsecase.DeleteLikebyThreadId(like[0])

		// assert.Equal(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})
}
