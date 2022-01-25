package bussiness_test

import (
	"errors"
	"os"
	"testing"

	"github.com/dragranzer/capstone-BE-FGD/features/comments"
	b_comments_mock "github.com/dragranzer/capstone-BE-FGD/features/comments/mocks"
	"github.com/dragranzer/capstone-BE-FGD/features/favorites"
	"github.com/dragranzer/capstone-BE-FGD/features/likes"
	b_likes_mock "github.com/dragranzer/capstone-BE-FGD/features/likes/mocks"
	"github.com/dragranzer/capstone-BE-FGD/features/threads"
	b_threads_mock "github.com/dragranzer/capstone-BE-FGD/features/threads/mocks"
	"github.com/dragranzer/capstone-BE-FGD/features/users"
	b_users_mock "github.com/dragranzer/capstone-BE-FGD/features/users/mocks"
	"github.com/stretchr/testify/assert"

	b_favorites "github.com/dragranzer/capstone-BE-FGD/features/favorites/bussiness"
	b_favorites_mock "github.com/dragranzer/capstone-BE-FGD/features/favorites/mocks"
)

var (
	commentUsecase  b_comments_mock.Bussiness
	userUsecase     b_users_mock.Bussiness
	threadUsecase   b_threads_mock.Bussiness
	likeUsecase     b_likes_mock.Bussiness
	favoriteData    b_favorites_mock.Data
	favoriteUsecase favorites.Bussiness

	comment  []comments.Core
	like     []likes.Core
	thread   []threads.Core
	user     []users.Core
	favorite []favorites.Core

	err1 error
)

func TestMain(m *testing.M) {
	favoriteUsecase = b_favorites.NewFavoriteBussiness(&threadUsecase, &userUsecase, &commentUsecase, &favoriteData, &likeUsecase)

	comment = []comments.Core{
		{
			// ID:       1,
			// Comment:  "tes",
			ThreadID: 1,
			// UserID:   1,
		},
	}

	like = []likes.Core{
		{
			// UserID:   1,
			ThreadID: 1,
		},
	}

	thread = []threads.Core{
		{
			ID: 1,
			// Title:       "judul 1",
			// Description: "desc1",
			// UserID: 1,
		},
	}

	user = []users.Core{
		{
			ID:       1,
			Email:    "email@email.com",
			Username: "saya",
		},
	}

	favorite = []favorites.Core{
		{
			// UserID:   1,
			ThreadID: 1,
		},
	}

	err1 = errors.New("tidak dapat menghapus komen orang lain")

	os.Exit(m.Run())
}

func TestAll(t *testing.T) {
	t.Run("valid - delete thread by id", func(t *testing.T) {
		commentUsecase.On("DeleteCommentbyThreadId", comment[0]).Return(nil).Once()
		likeUsecase.On("DeleteLikebyThreadId", like[0]).Return(nil).Once()
		favoriteData.On("DeleteFavoritebyThreadId", favorite[0]).Return(nil).Once()
		threadUsecase.On("DeleteThreadbyId", thread[0]).Return(nil).Once()
		err := favoriteUsecase.DeleteThreadbyId(favorite[0])

		// assert.Equal(t, resp, comment[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - delete thread by id 2", func(t *testing.T) {
		commentUsecase.On("DeleteCommentbyThreadId", comment[0]).Return(err1).Once()
		likeUsecase.On("DeleteLikebyThreadId", like[0]).Return(nil)
		favoriteData.On("DeleteFavoritebyThreadId", favorite[0]).Return(nil).Once()
		threadUsecase.On("DeleteThreadbyId", thread[0]).Return(nil).Once()
		err := favoriteUsecase.DeleteThreadbyId(favorite[0])

		// assert.Equal(t, resp, comment[0])
		assert.Equal(t, err, err1)
	})
}
