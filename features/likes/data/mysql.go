package data

import (
	"github.com/dragranzer/capstone-BE-FGD/features/likes"
	"gorm.io/gorm"
)

type mysqlLikeRepository struct {
	Conn *gorm.DB
}

func NewLikeRepository(conn *gorm.DB) likes.Data {
	return &mysqlLikeRepository{
		Conn: conn,
	}
}

func (fr *mysqlLikeRepository) InsertLike(data likes.Core) (err error) {
	recordData := fromCore(data)
	err = fr.Conn.Create(&recordData).Error
	if err != nil {
		return err
	}
	return err
}
