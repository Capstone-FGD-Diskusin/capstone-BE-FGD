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

func (ur *mysqlCategoryRepository) UpdateCategory(data categories.Core) (err error) {
	err = ur.Conn.Model(&Category{}).Where("id = ?", data.ID).Updates(map[string]interface{}{"name": data.Name}).Error
	if err != nil {
		return err
	}
	return err
}

func (ur *mysqlCategoryRepository) DeleteCategorybyId(data categories.Core) (err error) {
	record := Category{}
	err = ur.Conn.Where("id = ?", data.ID).Delete(&record).Error
	return
}

func (ur *mysqlCategoryRepository) SelectAllCategory(data categories.Core) (resp []categories.Core, err error) {
	record := []Category{}
	err = ur.Conn.Find(&record).Error
	resp = ToCoreSlice(record)
	return
}

func (ur *mysqlCategoryRepository) SelectCategorybyId(data categories.Core) (resp categories.Core, err error) {
	record := Category{}
	err = ur.Conn.Where("id = ?", data.ID).Find(&record).Error
	resp = record.toCore()
	return
}

func (ur *mysqlCategoryRepository) SelectCategorybyName(data categories.Core) (resp categories.Core, err error) {
	record := Category{}
	err = ur.Conn.Where("name = ?", data.Name).Find(&record).Error
	resp = record.toCore()
	return
}
