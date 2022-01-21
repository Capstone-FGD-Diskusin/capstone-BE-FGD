package data

import (
	"github.com/dragranzer/capstone-BE-FGD/features/messages"
	"gorm.io/gorm"
)

type mysqlMessageRepository struct {
	Conn *gorm.DB
}

func NewMessageRepository(conn *gorm.DB) messages.Data {
	return &mysqlMessageRepository{
		Conn: conn,
	}
}

func (mr *mysqlMessageRepository) InsertMessages(data messages.Core) (err error) {
	record := fromCore(data)
	err = mr.Conn.Create(&record).Error
	if err != nil {
		return err
	}
	return err
}

func (mr *mysqlMessageRepository) SelectMessagesbyAdminID(data messages.Core) (resp []messages.Core, err error) {
	record := []Message{}
	err = mr.Conn.Where("admin_id = ?", data.AdminID).Find(&record).Error
	if err != nil {
		return
	}
	resp = ToCoreSlice(record)
	return
}

func (mr *mysqlMessageRepository) DeleteMessagesbyId(data messages.Core) (err error) {
	record := Message{}
	err = mr.Conn.Where("id = ?", data.ID).Delete(&record).Error
	if err != nil {
		return
	}
	return
}
