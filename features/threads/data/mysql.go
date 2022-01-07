package data

import (
	"github.com/dragranzer/capstone-BE-FGD/features/threads"
	"gorm.io/gorm"
)

type mysqlThreadRepository struct {
	Conn *gorm.DB
}

func NewThreadRepository(conn *gorm.DB) threads.Data {
	return &mysqlThreadRepository{
		Conn: conn,
	}
}

func (ur *mysqlThreadRepository) SelectThreadHome(data threads.Core) (resp []threads.Core, err error) {

	return
}
