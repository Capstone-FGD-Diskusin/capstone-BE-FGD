package data

import (
	"github.com/dragranzer/capstone-BE-FGD/features/categories"
	"gorm.io/gorm"
)

type mysqlCategoryRepository struct {
	Conn *gorm.DB
}

func NewCategoryRepository(conn *gorm.DB) categories.Data {
	return &mysqlCategoryRepository{
		Conn: conn,
	}
}

func (ur *mysqlCategoryRepository) InsertCategory(data categories.Core) (err error) {
	record := Category{
		Name: data.Name,
	}
	err = ur.Conn.Create(&record).Error
	if err != nil {
		return err
	}
	return err
}
