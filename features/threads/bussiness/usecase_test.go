package bussiness_test

import (
	"errors"
	"os"
	"testing"

	"github.com/dragranzer/capstone-BE-FGD/features/categories"
	b_categories_mock "github.com/dragranzer/capstone-BE-FGD/features/categories/mocks"
	"github.com/dragranzer/capstone-BE-FGD/features/followers"
	b_followers_mock "github.com/dragranzer/capstone-BE-FGD/features/followers/mocks"
	"github.com/dragranzer/capstone-BE-FGD/features/threads"
	"github.com/dragranzer/capstone-BE-FGD/features/users"
	b_users_mock "github.com/dragranzer/capstone-BE-FGD/features/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	b_threads "github.com/dragranzer/capstone-BE-FGD/features/threads/bussiness"
	b_threads_mock "github.com/dragranzer/capstone-BE-FGD/features/threads/mocks"
)

var (
	followerUsecase b_followers_mock.Bussiness
	userUsecase     b_users_mock.Bussiness
	categoryUsecase b_categories_mock.Bussiness
	threadData      b_threads_mock.Data
	threadUsecase   threads.Bussiness

	thread   []threads.Core
	user     []users.Core
	category []categories.Core
	follower []followers.Core

	err1 error
)

func TestMain(m *testing.M) {
	threadUsecase = b_threads.NewThreadBussiness(&followerUsecase, &threadData, &categoryUsecase, &userUsecase)

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

	follower = []followers.Core{
		{
			FollowingID: 1,
			FollowedID:  2,
		},
	}

	category = []categories.Core{
		{
			ID:   1,
			Name: "Hiburan",
		},
	}

	err1 = errors.New("tidak dapat menghapus komen orang lain")

	os.Exit(m.Run())
}

func TestAll(t *testing.T) {
	t.Run("valid - get thread home", func(t *testing.T) {
		followerUsecase.On("GetFollowing", mock.AnythingOfType("followers.Core")).Return(follower, nil).Once()
		threadData.On("SelectThreadHome", mock.AnythingOfType("threads.Core")).Return(thread, nil).Once()
		categoryUsecase.On("GetCategorybyId", mock.AnythingOfType("categories.Core")).Return(category[0], nil).Once()
		userUsecase.On("GetProfileData", mock.AnythingOfType("users.Core")).Return(user[0], nil).Once()
		resp, err := threadUsecase.GetThreadHome(thread[0])

		assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - get thread home", func(t *testing.T) {
		followerUsecase.On("GetFollowing", mock.AnythingOfType("followers.Core")).Return(follower, err1).Once()
		threadData.On("SelectThreadHome", mock.AnythingOfType("threads.Core")).Return(thread, nil).Once()
		categoryUsecase.On("GetCategorybyId", mock.AnythingOfType("categories.Core")).Return(category[0], nil).Once()
		userUsecase.On("GetProfileData", mock.AnythingOfType("users.Core")).Return(user[0], nil).Once()
		resp, err := threadUsecase.GetThreadHome(thread[0])

		assert.Equal(t, len(resp), 0)
		assert.NotEqual(t, err, nil)
	})

	t.Run("valid - add thread", func(t *testing.T) {
		categoryUsecase.On("GetCategorybyName", mock.AnythingOfType("categories.Core")).Return(category[0], nil).Once()
		threadData.On("InsertThread", mock.AnythingOfType("threads.Core")).Return(nil).Once()
		userUsecase.On("IncrementThread", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := threadUsecase.AddThread(thread[0])

		// assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - add thread", func(t *testing.T) {
		categoryUsecase.On("GetCategorybyName", mock.AnythingOfType("categories.Core")).Return(category[0], err1).Once()
		threadData.On("InsertThread", mock.AnythingOfType("threads.Core")).Return(nil).Once()
		userUsecase.On("IncrementThread", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := threadUsecase.AddThread(thread[0])

		// assert.NotEqual(t, len(resp), 0)
		assert.NotEqual(t, err, nil)
	})

	t.Run("valid - get thread by id", func(t *testing.T) {
		threadData.On("SelectThreadbyID", mock.AnythingOfType("threads.Core")).Return(thread[0], nil).Once()
		categoryUsecase.On("GetCategorybyId", mock.AnythingOfType("categories.Core")).Return(category[0], nil).Once()
		userUsecase.On("GetProfileData", mock.AnythingOfType("users.Core")).Return(user[0], nil).Once()
		resp, err := threadUsecase.GetThreadbyID(thread[0])

		assert.Equal(t, resp, thread[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - get thread by id", func(t *testing.T) {
		threadData.On("SelectThreadbyID", mock.AnythingOfType("threads.Core")).Return(thread[0], err1).Once()
		categoryUsecase.On("GetCategorybyId", mock.AnythingOfType("categories.Core")).Return(category[0], nil).Once()
		userUsecase.On("GetProfileData", mock.AnythingOfType("users.Core")).Return(user[0], nil).Once()
		resp, err := threadUsecase.GetThreadbyID(thread[0])

		assert.Equal(t, resp, thread[0])
		assert.NotEqual(t, err, nil)
	})

	t.Run("valid - increment like", func(t *testing.T) {
		threadData.On("UpdateLikebyOne", mock.AnythingOfType("threads.Core")).Return(nil).Once()
		err := threadUsecase.IncrementLike(thread[0])

		// assert.Equal(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - decrement like", func(t *testing.T) {
		threadData.On("UpdateMinLikebyOne", mock.AnythingOfType("threads.Core")).Return(nil).Once()
		err := threadUsecase.DecrementLike(thread[0])

		// assert.Equal(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - increment comment", func(t *testing.T) {
		threadData.On("UpdateCommentbyOne", mock.AnythingOfType("threads.Core")).Return(nil).Once()
		err := threadUsecase.IncrementComment(thread[0])

		// assert.Equal(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - get thread by id", func(t *testing.T) {
		threadData.On("SelectThreadbyID", mock.AnythingOfType("threads.Core")).Return(thread[0], nil).Once()
		threadData.On("DeleteThreadbyId", mock.AnythingOfType("threads.Core")).Return(nil).Once()
		userUsecase.On("DecrementThread", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := threadUsecase.DeleteThreadbyId(thread[0])

		// assert.Equal(t, resp, thread[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - get thread by id", func(t *testing.T) {
		threadData.On("SelectThreadbyID", mock.AnythingOfType("threads.Core")).Return(thread[0], err1).Once()
		threadData.On("DeleteThreadbyId", mock.AnythingOfType("threads.Core")).Return(nil).Once()
		userUsecase.On("DecrementThread", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := threadUsecase.DeleteThreadbyId(thread[0])

		// assert.Equal(t, resp, thread[0])
		assert.NotEqual(t, err, nil)
	})

	t.Run("valid - search thread", func(t *testing.T) {
		threadData.On("SearchThread", mock.AnythingOfType("threads.Core")).Return(thread, nil).Once()
		categoryUsecase.On("GetCategorybyId", mock.AnythingOfType("categories.Core")).Return(category[0], nil).Once()
		userUsecase.On("GetProfileData", mock.AnythingOfType("users.Core")).Return(user[0], nil).Once()
		resp, err := threadUsecase.SearchThread(thread[0])

		assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - search thread", func(t *testing.T) {
		threadData.On("SearchThread", mock.AnythingOfType("threads.Core")).Return(thread, err1).Once()
		categoryUsecase.On("GetCategorybyId", mock.AnythingOfType("categories.Core")).Return(category[0], nil).Once()
		userUsecase.On("GetProfileData", mock.AnythingOfType("users.Core")).Return(user[0], nil).Once()
		resp, err := threadUsecase.SearchThread(thread[0])

		assert.NotEqual(t, len(resp), 0)
		assert.NotEqual(t, err, nil)
	})

	t.Run("valid - get thread all", func(t *testing.T) {
		threadData.On("SelectThreadAll", mock.AnythingOfType("threads.Core")).Return(thread, nil).Once()
		categoryUsecase.On("GetCategorybyId", mock.AnythingOfType("categories.Core")).Return(category[0], nil).Once()
		userUsecase.On("GetProfileData", mock.AnythingOfType("users.Core")).Return(user[0], nil).Once()
		resp, err := threadUsecase.GetAllThread(thread[0])

		assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - get thread all", func(t *testing.T) {
		threadData.On("SelectThreadAll", mock.AnythingOfType("threads.Core")).Return(thread, err1).Once()
		categoryUsecase.On("GetCategorybyId", mock.AnythingOfType("categories.Core")).Return(category[0], nil).Once()
		userUsecase.On("GetProfileData", mock.AnythingOfType("users.Core")).Return(user[0], nil).Once()
		resp, err := threadUsecase.GetAllThread(thread[0])

		assert.NotEqual(t, len(resp), 0)
		assert.NotEqual(t, err, nil)
	})

	t.Run("valid - get thread user", func(t *testing.T) {
		threadData.On("SelectThreadUser", mock.AnythingOfType("threads.Core")).Return(thread, nil).Once()
		categoryUsecase.On("GetCategorybyId", mock.AnythingOfType("categories.Core")).Return(category[0], nil).Once()
		userUsecase.On("GetProfileData", mock.AnythingOfType("users.Core")).Return(user[0], nil).Once()
		resp, err := threadUsecase.GetThreadUser(thread[0])

		assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - get thread user", func(t *testing.T) {
		threadData.On("SelectThreadUser", mock.AnythingOfType("threads.Core")).Return(thread, err1).Once()
		categoryUsecase.On("GetCategorybyId", mock.AnythingOfType("categories.Core")).Return(category[0], nil).Once()
		userUsecase.On("GetProfileData", mock.AnythingOfType("users.Core")).Return(user[0], nil).Once()
		resp, err := threadUsecase.GetThreadUser(thread[0])

		assert.NotEqual(t, len(resp), 0)
		assert.NotEqual(t, err, nil)
	})
}
