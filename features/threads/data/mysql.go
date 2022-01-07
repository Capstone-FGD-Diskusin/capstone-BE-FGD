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

func (tr *mysqlThreadRepository) SelectThreadHome(data threads.Core) (resp []threads.Core, err error) {
	// fmt.Println(data)
	record := []Thread{}
	err = tr.Conn.Limit(20).Offset(data.Page*20).Where("user_id IN ?", data.ListFollowedID).Find(&record).Error
	// fmt.Println(record)
	resp = ToCoreSlice(record)
	return
}
