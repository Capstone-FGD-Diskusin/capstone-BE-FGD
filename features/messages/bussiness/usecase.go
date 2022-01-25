package bussiness

import (
	"errors"

	"github.com/dragranzer/capstone-BE-FGD/features/comments"
	"github.com/dragranzer/capstone-BE-FGD/features/messages"
	"github.com/dragranzer/capstone-BE-FGD/features/threads"
	"github.com/dragranzer/capstone-BE-FGD/features/users"
)

type messagesUsecase struct {
	messageData      messages.Data
	userBussiness    users.Bussiness
	threadBussiness  threads.Bussiness
	commentBussiness comments.Bussiness
}

func NewMessageBussiness(messageData messages.Data, ub users.Bussiness, tB threads.Bussiness, cB comments.Bussiness) messages.Bussiness {
	return &messagesUsecase{
		messageData:      messageData,
		userBussiness:    ub,
		threadBussiness:  tB,
		commentBussiness: cB,
	}
}

func (mu *messagesUsecase) ReportToAdmin(data messages.Core) (err error) {
	userCore := users.Core{
		ID: data.ModeratorID,
	}
	user, err := mu.userBussiness.GetProfileData(userCore)
	if err != nil {
		return
	}
	if user.Role != "moderator" {
		err = errors.New("bukan moderator jangan sok keras :v")
	}
	if err != nil {
		return
	}
	data.ModeratorName = user.Username
	err = mu.messageData.InsertMessages(data)
	return err
}

func (mu *messagesUsecase) GetMessagesbyAdminID(data messages.Core) (resp []messages.Core, err error) {
	userCore := users.Core{
		ID: data.AdminID,
	}
	user, err := mu.userBussiness.GetProfileData(userCore)
	if err != nil {
		return
	}
	if user.Role != "admin" {
		err = errors.New("jadi admin dulu yaaa :v")
	}
	if err != nil {
		return
	}
	resp, err = mu.messageData.SelectMessagesbyAdminID(data)
	for key, value := range resp {
		threadCore := threads.Core{
			ID: value.ThreadID,
		}
		thread, err := mu.threadBussiness.GetThreadbyID(threadCore)
		if err != nil {
			continue
		}
		commentCore := comments.Core{
			ID: value.CommentID,
		}
		comment, err := mu.commentBussiness.GetCommentbyId(commentCore)
		if err != nil {
			continue
		}
		resp[key].ThreadTitle = thread.Title
		resp[key].AdminName = user.Username
		resp[key].Comment = comment.Comment
	}

	return
}

func (mu *messagesUsecase) DeleteMessagesbyId(data messages.Core) (err error) {
	userCore := users.Core{
		ID: data.AdminID,
	}
	user, err := mu.userBussiness.GetProfileData(userCore)
	if err != nil {
		return
	}
	if user.Role != "admin" {
		err = errors.New("jadi admin dulu yaaa :v")
	}
	if err != nil {
		return
	}
	err = mu.messageData.DeleteMessagesbyId(data)
	return
}
