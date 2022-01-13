package data

import (
	"github.com/dragranzer/capstone-BE-FGD/features/comments"
	"gorm.io/gorm"
)

type mysqlCommentRepository struct {
	Conn *gorm.DB
}

func NewCommentRepository(conn *gorm.DB) comments.Data {
	return &mysqlCommentRepository{
		Conn: conn,
	}
}

func (cr *mysqlCommentRepository) InsertComment(data comments.Core) (err error) {
	record := FromCore(data)
	err = cr.Conn.Create(&record).Error
	return
}

func (cr *mysqlCommentRepository) SelectCommentsThread(data comments.Core) (resp []comments.Core, err error) {
	record := []Comment{}
	err = cr.Conn.Where("thread_id = ?", data.ThreadID).Find(&record).Error
	if err != nil {
		return resp, err
	}
	for _, value := range record {
		resp = append(resp, ToCore(value))
	}
	return
}
