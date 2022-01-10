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

	return
}
