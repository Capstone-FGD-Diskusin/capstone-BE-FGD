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

func (ur *mysqlUserRepository) CheckEmailPass(email string, pass string) (isAuth bool, user users.Core, err error) {
	record := User{}
	err = ur.Conn.Where("email = ? AND password = ?", email, pass).First(&record).Error
	if err != nil {
		return false, user, err
	}
	if record.ID == 0 {
		return false, user, nil
	}
	return true, record.toCore(), nil
}
