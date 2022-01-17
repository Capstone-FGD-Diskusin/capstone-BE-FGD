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
