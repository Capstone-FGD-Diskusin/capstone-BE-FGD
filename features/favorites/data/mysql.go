package data

import (
	"github.com/dragranzer/capstone-BE-FGD/features/favorites"
	"gorm.io/gorm"
)

type mysqlFavoriteRepository struct {
	Conn *gorm.DB
}

func NewFavoriteRepository(conn *gorm.DB) favorites.Data {
	return &mysqlFavoriteRepository{
		Conn: conn,
	}
}

func (cr *mysqlFavoriteRepository) DeleteFavoritebyThreadId(data favorites.Core) (err error) {
	record := Favorite{}
	err = cr.Conn.Where("thread_id = ?", data.ThreadID).Delete(&record).Error
	return
}

func (cr *mysqlFavoriteRepository) AddFavorite(data favorites.Core) (err error) {
	record := FromCore(data)
	err = cr.Conn.Create(&record).Error
	return
}

func (cr *mysqlFavoriteRepository) DeleteFavorite(data favorites.Core) (err error) {
	record := Favorite{}
	err = cr.Conn.Where("thread_id = ? AND user_id = ?", data.ThreadID, data.UserID).Delete(&record).Error
	return
}
