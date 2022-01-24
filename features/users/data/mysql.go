package data

import (
	"github.com/dragranzer/capstone-BE-FGD/features/users"
	"golang.org/x/crypto/bcrypt"
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
	bytes, _ := bcrypt.GenerateFromPassword([]byte(recordData.Password), 14)
	recordData.Password = string(bytes)
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

func (ur *mysqlUserRepository) SelectDatabyEmail(data users.Core) (resp users.Core, err error) {
	record := User{}
	if err = ur.Conn.Where("email = ?", data.Email).Find(&record).Error; err != nil {
		return users.Core{}, err
	}

	return record.toCore(), nil
}

func (ur *mysqlUserRepository) SelectDatabyID(data users.Core) (resp users.Core, err error) {
	record := User{}
	if err = ur.Conn.Where("id = ?", data.ID).Find(&record).Error; err != nil {
		return users.Core{}, err
	}
	return record.toCore(), nil
}

func (ur *mysqlUserRepository) UpdateLikebyOne(data users.Core) (err error) {
	record := User{}
	err = ur.Conn.Where("id = ?", data.ID).First(&record).Error
	if err != nil {
		return err
	}
	record.SumLike++
	err = ur.Conn.Model(&User{}).Where("id = ?", data.ID).Update("sum_like", record.SumLike).Error
	return
}

func (ur *mysqlUserRepository) UpdateMinLikebyOne(data users.Core) (err error) {
	record := User{}
	err = ur.Conn.Where("id = ?", data.ID).First(&record).Error
	if err != nil {
		return err
	}
	record.SumLike--
	err = ur.Conn.Model(&User{}).Where("id = ?", data.ID).Update("sum_like", record.SumLike).Error
	return
}

func (ur *mysqlUserRepository) UpdateThreadbyOne(data users.Core) (err error) {
	record := User{}
	err = ur.Conn.Where("id = ?", data.ID).First(&record).Error
	if err != nil {
		return err
	}
	record.SumThread++
	err = ur.Conn.Model(&User{}).Where("id = ?", data.ID).Update("sum_thread", record.SumThread).Error
	return
}

func (ur *mysqlUserRepository) UpdateMinThreadbyOne(data users.Core) (err error) {
	record := User{}
	err = ur.Conn.Where("id = ?", data.ID).First(&record).Error
	if err != nil {
		return err
	}
	record.SumThread--
	err = ur.Conn.Model(&User{}).Where("id = ?", data.ID).Update("sum_thread", record.SumThread).Error
	return
}

func (ur *mysqlUserRepository) UpdateFolbyOne(data users.Core) (err error) {
	record := User{}
	err = ur.Conn.Where("id = ?", data.ID).First(&record).Error
	if err != nil {
		return err
	}
	record.Follower++
	err = ur.Conn.Model(&User{}).Where("id = ?", data.ID).Update("follower", record.Follower).Error
	return
}

func (ur *mysqlUserRepository) UpdateMinFolbyOne(data users.Core) (err error) {
	record := User{}
	err = ur.Conn.Where("id = ?", data.ID).First(&record).Error
	if err != nil {
		return err
	}
	record.Follower--
	err = ur.Conn.Model(&User{}).Where("id = ?", data.ID).Update("follower", record.Follower).Error
	return
}

func (ur *mysqlUserRepository) UpdateFollowingbyOne(data users.Core) (err error) {
	record := User{}
	err = ur.Conn.Where("id = ?", data.ID).First(&record).Error
	if err != nil {
		return err
	}
	record.Following++
	err = ur.Conn.Model(&User{}).Where("id = ?", data.ID).Update("following", record.Following).Error
	return
}

func (ur *mysqlUserRepository) UpdateMinFollowingbyOne(data users.Core) (err error) {
	record := User{}
	err = ur.Conn.Where("id = ?", data.ID).First(&record).Error
	if err != nil {
		return err
	}
	record.Following--
	err = ur.Conn.Model(&User{}).Where("id = ?", data.ID).Update("following", record.Following).Error
	return
}

func (ur *mysqlUserRepository) UpdateDataUser(data users.Core) (err error) {
	record := fromCore(data)
	bytes, _ := bcrypt.GenerateFromPassword([]byte(record.Password), 14)
	record.Password = string(bytes)
	err = ur.Conn.Model(&User{}).Where("id = ?", data.ID).Updates(map[string]interface{}{"email": record.Email, "username": record.Username, "profile_picture": record.ProfilePicture}).Error
	return
}

func (ur *mysqlUserRepository) DeleteDataUserbyId(data users.Core) (err error) {
	record := User{}
	err = ur.Conn.Where("id = ?", data.ID).Delete(&record).Error
	return
}

func (ur *mysqlUserRepository) UpdateUserToModerator(data users.Core) (err error) {
	err = ur.Conn.Model(&User{}).Where("id = ?", data.ID).Update("category_id", data.CategoryID).Error
	if err != nil {
		return err
	}
	err = ur.Conn.Model(&User{}).Where("id = ?", data.ID).Update("role", "moderator").Error
	return
}

func (ur *mysqlUserRepository) SelectAllUser(data users.Core) (resp []users.Core, err error) {
	record := []User{}
	err = ur.Conn.Limit(20).Offset(data.Page * 20).Find(&record).Error
	resp = ToCoreSlice(record)
	return
}

func (ur *mysqlUserRepository) Ranking() (resp []users.Core, err error) {
	record := []User{}
	err = ur.Conn.Limit(10).Order("sum_like desc").Find(&record).Error
	resp = ToCoreSlice(record)
	return
}
