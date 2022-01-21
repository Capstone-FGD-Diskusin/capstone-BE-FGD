package bussiness

import (
	"errors"

	"github.com/dragranzer/capstone-BE-FGD/features/messages"
	"github.com/dragranzer/capstone-BE-FGD/features/users"
)

type messagesUsecase struct {
	messageData   messages.Data
	userBussiness users.Bussiness
}

func NewMessageBussiness(messageData messages.Data, ub users.Bussiness) messages.Bussiness {
	return &messagesUsecase{
		messageData:   messageData,
		userBussiness: ub,
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
	err = mu.messageData.InsertMessages(data)
	return err
}

func (mu *messagesUsecase) GetMessagesbyAdminID(data messages.Core) (resp []messages.Core, err error) {
	err = mu.messageData.InsertMessages(data)
	return
}

func (mu *messagesUsecase) DeleteMessagesbyId(data messages.Core) (err error) {
	err = mu.messageData.InsertMessages(data)
	return
}
