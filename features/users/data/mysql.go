package data

import (
	"github.com/dragranzer/capstone-BE-FGD/features/users"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) users.Data {
	return &mysqlUserRepository{
		Conn: conn,
	}
}

func (ur *mysqlUserRepository) CreateUser(data users.Core) (err error) {
	recordData := fromCore(data)
	err = ur.Conn.Create(&recordData).Error
	if err != nil {
		return err
	}
	return err
}
