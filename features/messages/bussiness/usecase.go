package bussiness

import "github.com/dragranzer/capstone-BE-FGD/features/messages"

type messagesUsecase struct {
	messageData messages.Data
}

func NewMessageBussiness(messageData messages.Data) messages.Bussiness {
	return &messagesUsecase{
		messageData: messageData,
	}
}

func (mu *messagesUsecase) ReportToAdmin(data messages.Core) (err error) {
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
