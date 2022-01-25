package bussiness_test

import (
	"errors"
	"os"
	"testing"

	"github.com/dragranzer/capstone-BE-FGD/features/users"
	b_users "github.com/dragranzer/capstone-BE-FGD/features/users/bussiness"
	b_users_mock "github.com/dragranzer/capstone-BE-FGD/features/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userData    b_users_mock.Data
	userUsecase users.Bussiness

	user []users.Core

	err1 error
)

func TestMain(m *testing.M) {
	userUsecase = b_users.NewUserBussiness(&userData)

	user = []users.Core{
		{
			ID:       1,
			Email:    "email@email.com",
			Username: "saya",
			Role:     "user",
		},
		{
			ID:       2,
			Email:    "email@email.com",
			Username: "saya",
			Role:     "admin",
		},
	}

	err1 = errors.New("tidak dapat menghapus komen orang lain")

	os.Exit(m.Run())
}

func TestAll(t *testing.T) {
	t.Run("valid - register", func(t *testing.T) {
		userData.On("CreateUser", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := userUsecase.Register(user[0])

		// assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - login", func(t *testing.T) {
		userData.On("SelectDatabyEmail", mock.AnythingOfType("users.Core")).Return(user[0], nil).Once()
		token := ""
		isAuth := true
		resp, token, isAuth, err := userUsecase.Login(user[0])

		assert.Equal(t, resp, user[0])
		assert.Equal(t, err, nil)
		assert.NotEqual(t, token, "")
		assert.Equal(t, isAuth, false)
	})

	t.Run("valid - get profile data", func(t *testing.T) {
		userData.On("SelectDatabyID", mock.AnythingOfType("users.Core")).Return(user[0], nil).Once()
		resp, err := userUsecase.GetProfileData(user[0])

		assert.Equal(t, resp, user[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - increment like", func(t *testing.T) {
		userData.On("UpdateLikebyOne", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := userUsecase.IncrementLike(user[0])

		// assert.Equal(t, resp, user[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - decrement like", func(t *testing.T) {
		userData.On("UpdateMinLikebyOne", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := userUsecase.DecrementLike(user[0])

		// assert.Equal(t, resp, user[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - increment thread", func(t *testing.T) {
		userData.On("UpdateThreadbyOne", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := userUsecase.IncrementThread(user[0])

		// assert.Equal(t, resp, user[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - decrement thread", func(t *testing.T) {
		userData.On("UpdateMinThreadbyOne", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := userUsecase.DecrementThread(user[0])

		// assert.Equal(t, resp, user[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - increment fol", func(t *testing.T) {
		userData.On("UpdateFolbyOne", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := userUsecase.IncrementFol(user[0])

		// assert.Equal(t, resp, user[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - decrement fol", func(t *testing.T) {
		userData.On("UpdateMinFolbyOne", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := userUsecase.DecrementFol(user[0])

		// assert.Equal(t, resp, user[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - increment following", func(t *testing.T) {
		userData.On("UpdateFollowingbyOne", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := userUsecase.IncrementFollowing(user[0])

		// assert.Equal(t, resp, user[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - decrement following", func(t *testing.T) {
		userData.On("UpdateMinFollowingbyOne", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := userUsecase.DecrementFollowing(user[0])

		// assert.Equal(t, resp, user[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - edit data user", func(t *testing.T) {
		userData.On("UpdateDataUser", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := userUsecase.EditDataUser(user[0])

		// assert.Equal(t, resp, user[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - delete data user by id", func(t *testing.T) {
		userData.On("DeleteDataUserbyId", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := userUsecase.DeleteDataUserbyId(user[0])

		// assert.Equal(t, resp, user[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - upgrade to moderator", func(t *testing.T) {
		userData.On("SelectDatabyID", mock.AnythingOfType("users.Core")).Return(user[0], nil).Once()
		userData.On("UpdateUserToModerator", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := userUsecase.UpgradeToModerator(user[0])

		// assert.Equal(t, resp, user[0])
		assert.NotEqual(t, err, nil)
	})

	t.Run("valid - upgrade to moderator", func(t *testing.T) {
		userData.On("SelectDatabyID", mock.AnythingOfType("users.Core")).Return(user[1], nil).Once()
		userData.On("UpdateUserToModerator", mock.AnythingOfType("users.Core")).Return(nil).Once()
		err := userUsecase.UpgradeToModerator(user[0])

		// assert.Equal(t, resp, user[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - get all user", func(t *testing.T) {
		userData.On("SelectAllUser", mock.AnythingOfType("users.Core")).Return(user, nil).Once()
		resp, err := userUsecase.GetAllUser(user[0])

		assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - ranking", func(t *testing.T) {
		userData.On("Ranking").Return(user, nil).Once()
		resp, err := userUsecase.Ranking()

		assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})
}
