package bussiness_test

import (
	"errors"
	"os"
	"testing"

	"github.com/dragranzer/capstone-BE-FGD/features/followers"
	b_followers "github.com/dragranzer/capstone-BE-FGD/features/followers/bussiness"
	b_followers_mock "github.com/dragranzer/capstone-BE-FGD/features/followers/mocks"
	"github.com/dragranzer/capstone-BE-FGD/features/users"
	b_users_mock "github.com/dragranzer/capstone-BE-FGD/features/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userUsecase     b_users_mock.Bussiness
	followerData    b_followers_mock.Data
	followerUsecase followers.Bussiness

	follower []followers.Core
	user     []users.Core

	err1 error
)

func TestMain(m *testing.M) {
	followerUsecase = b_followers.NewFollowerBussiness(&followerData, &userUsecase)
	follower = []followers.Core{
		{
			FollowingID: 1,
			FollowedID:  2,
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
	t.Run("valid - follow", func(t *testing.T) {
		followerData.On("InsertFollow", mock.AnythingOfType("followers.Core")).Return(nil).Once()
		userUsecase.On("IncrementFol", mock.AnythingOfType("users.Core")).Return(nil).Once()
		userUsecase.On("IncrementFollowing", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := followerUsecase.Follow(follower[0])

		// assert.Equal(t, resp, comment[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - follow", func(t *testing.T) {
		followerData.On("InsertFollow", mock.AnythingOfType("followers.Core")).Return(err1).Once()
		userUsecase.On("IncrementFol", mock.AnythingOfType("users.Core")).Return(nil).Once()
		userUsecase.On("IncrementFollowing", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := followerUsecase.Follow(follower[0])

		// assert.Equal(t, resp, comment[0])
		assert.Equal(t, err, err1)
	})

	t.Run("valid - unfollow", func(t *testing.T) {
		followerData.On("DeleteFollow", mock.AnythingOfType("followers.Core")).Return(nil).Once()
		userUsecase.On("DecrementFol", mock.AnythingOfType("users.Core")).Return(nil).Once()
		userUsecase.On("DecrementFollowing", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := followerUsecase.Unfollow(follower[0])

		// assert.Equal(t, resp, comment[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - unfollow", func(t *testing.T) {
		followerData.On("DeleteFollow", mock.AnythingOfType("followers.Core")).Return(err1).Once()
		userUsecase.On("DecrementFol", mock.AnythingOfType("users.Core")).Return(nil).Once()
		userUsecase.On("DecrementFollowing", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := followerUsecase.Unfollow(follower[0])

		// assert.Equal(t, resp, comment[0])
		assert.Equal(t, err, err1)
	})

	t.Run("valid - get following", func(t *testing.T) {
		followerData.On("SelectFollowing", mock.AnythingOfType("followers.Core")).Return(follower, nil).Once()
		userUsecase.On("GetProfileData", mock.AnythingOfType("users.Core")).Return(user[0], nil).Once()

		resp, err := followerUsecase.GetFollowing(follower[0])

		assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - get followed", func(t *testing.T) {
		followerData.On("SelectFollowed", mock.AnythingOfType("followers.Core")).Return(follower, nil).Once()
		userUsecase.On("GetProfileData", mock.AnythingOfType("users.Core")).Return(user[0], nil).Once()

		resp, err := followerUsecase.GetFollowed(follower[0])

		assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})
}
