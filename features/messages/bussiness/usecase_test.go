package bussiness_test

import (
	"errors"
	"os"
	"testing"

	"github.com/dragranzer/capstone-BE-FGD/features/comments"
	b_comments_mock "github.com/dragranzer/capstone-BE-FGD/features/comments/mocks"
	"github.com/dragranzer/capstone-BE-FGD/features/messages"
	"github.com/dragranzer/capstone-BE-FGD/features/threads"
	b_threads_mock "github.com/dragranzer/capstone-BE-FGD/features/threads/mocks"
	"github.com/dragranzer/capstone-BE-FGD/features/users"
	b_users_mock "github.com/dragranzer/capstone-BE-FGD/features/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	b_messages "github.com/dragranzer/capstone-BE-FGD/features/messages/bussiness"
	b_messages_mock "github.com/dragranzer/capstone-BE-FGD/features/messages/mocks"
)

var (
	userUsecase    b_users_mock.Bussiness
	threadUsecase  b_threads_mock.Bussiness
	commentUsecase b_comments_mock.Bussiness
	messageData    b_messages_mock.Data
	messageUsecase messages.Bussiness

	thread  []threads.Core
	user    []users.Core
	message []messages.Core
	comment []comments.Core

	err1 error
)

func TestMain(m *testing.M) {
	messageUsecase = b_messages.NewMessageBussiness(&messageData, &userUsecase, &threadUsecase, &commentUsecase)

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
			Role:     "moderator",
		},
		{
			ID:       2,
			Email:    "email@email.com",
			Username: "saya",
			Role:     "user",
		},
		{
			ID:       3,
			Email:    "email@email.com",
			Username: "saya",
			Role:     "admin",
		},
	}

	comment = []comments.Core{
		{
			ID:       1,
			Comment:  "tes",
			ThreadID: 1,
			UserID:   1,
		},
	}

	message = []messages.Core{
		{
			ID:       1,
			Comment:  "tes",
			ThreadID: 1,
			AdminID:  1,
		},
	}
	err1 = errors.New("tidak dapat menghapus komen orang lain")

	os.Exit(m.Run())
}

func TestAll(t *testing.T) {
	t.Run("valid - report to admin", func(t *testing.T) {
		userUsecase.On("GetProfileData", mock.AnythingOfType("users.Core")).Return(user[0], nil).Once()
		messageData.On("InsertMessages", mock.AnythingOfType("messages.Core")).Return(nil).Once()
		err := messageUsecase.ReportToAdmin(message[0])

		// assert.Equal(t, resp, comment[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - report to admin", func(t *testing.T) {
		userUsecase.On("GetProfileData", mock.AnythingOfType("users.Core")).Return(user[1], nil).Once()
		messageData.On("InsertMessages", mock.AnythingOfType("messages.Core")).Return(nil).Once()
		err := messageUsecase.ReportToAdmin(message[0])

		// assert.Equal(t, resp, comment[0])
		assert.NotEqual(t, err, nil)
	})

	t.Run("valid - report to admin", func(t *testing.T) {
		userUsecase.On("GetProfileData", mock.AnythingOfType("users.Core")).Return(user[1], err1).Once()
		messageData.On("InsertMessages", mock.AnythingOfType("messages.Core")).Return(nil).Once()
		err := messageUsecase.ReportToAdmin(message[0])

		// assert.Equal(t, resp, comment[0])
		assert.NotEqual(t, err, nil)
	})

	t.Run("valid - get message by admin id", func(t *testing.T) {
		userUsecase.On("GetProfileData", mock.AnythingOfType("users.Core")).Return(user[2], nil).Once()
		messageData.On("SelectMessagesbyAdminID", mock.AnythingOfType("messages.Core")).Return(message, nil).Once()
		threadUsecase.On("GetThreadbyID", mock.AnythingOfType("threads.Core")).Return(thread[0], nil).Once()
		commentUsecase.On("GetCommentbyId", mock.AnythingOfType("comments.Core")).Return(comment[0], nil).Once()
		resp, err := messageUsecase.GetMessagesbyAdminID(message[0])

		assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})

	t.Run("valid - get message by admin id", func(t *testing.T) {
		userUsecase.On("GetProfileData", mock.AnythingOfType("users.Core")).Return(user[0], nil).Once()
		messageData.On("SelectMessagesbyAdminID", mock.AnythingOfType("messages.Core")).Return(message, nil).Once()
		threadUsecase.On("GetThreadbyID", mock.AnythingOfType("threads.Core")).Return(thread[0], nil).Once()
		commentUsecase.On("GetCommentbyId", mock.AnythingOfType("comments.Core")).Return(comment[0], nil).Once()
		resp, err := messageUsecase.GetMessagesbyAdminID(message[0])

		assert.Equal(t, len(resp), 0)
		assert.NotEqual(t, err, nil)
	})

	t.Run("valid - delete message by id", func(t *testing.T) {
		userUsecase.On("GetProfileData", mock.AnythingOfType("users.Core")).Return(user[2], nil).Once()
		messageData.On("DeleteMessagesbyId", mock.AnythingOfType("messages.Core")).Return(nil).Once()
		err := messageUsecase.DeleteMessagesbyId(message[0])

		// assert.Equal(t, resp, comment[0])
		assert.Equal(t, err, nil)
	})

	t.Run("valid - delete message by id", func(t *testing.T) {
		userUsecase.On("GetProfileData", mock.AnythingOfType("users.Core")).Return(user[1], nil).Once()
		messageData.On("DeleteMessagesbyId", mock.AnythingOfType("messages.Core")).Return(nil).Once()
		err := messageUsecase.DeleteMessagesbyId(message[0])

		// assert.Equal(t, resp, comment[0])
		assert.NotEqual(t, err, nil)
	})

	t.Run("valid - delete message by id", func(t *testing.T) {
		userUsecase.On("GetProfileData", mock.AnythingOfType("users.Core")).Return(user[2], err1).Once()
		messageData.On("DeleteMessagesbyId", mock.AnythingOfType("messages.Core")).Return(nil).Once()
		err := messageUsecase.DeleteMessagesbyId(message[0])

		// assert.Equal(t, resp, comment[0])
		assert.NotEqual(t, err, nil)
	})
}
