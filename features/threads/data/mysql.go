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

func (tr *mysqlThreadRepository) InsertThread(data threads.Core) (err error) {
	recordData := FromCore(data)
	err = tr.Conn.Create(&recordData).Error
	return
}

func (tr *mysqlThreadRepository) SelectThreadbyID(data threads.Core) (resp threads.Core, err error) {
	record := Thread{}
	err = tr.Conn.Where("id = ?", data.ID).First(&record).Error
	resp = ToCore(record)
	return
}

func (ur *mysqlThreadRepository) UpdateLikebyOne(data threads.Core) (err error) {
	record := Thread{}
	err = ur.Conn.Where("id = ?", data.ID).First(&record).Error
	if err != nil {
		return err
	}
	record.Like++
	err = ur.Conn.Model(&Thread{}).Where("id = ?", data.ID).Update("like", record.Like).Error
	return
}

func (ur *mysqlThreadRepository) UpdateMinLikebyOne(data threads.Core) (err error) {
	record := Thread{}
	err = ur.Conn.Where("id = ?", data.ID).First(&record).Error
	if err != nil {
		return err
	}
	record.Like--
	err = ur.Conn.Model(&Thread{}).Where("id = ?", data.ID).Update("like", record.Like).Error
	return
}

func (ur *mysqlThreadRepository) UpdateCommentbyOne(data threads.Core) (err error) {
	record := Thread{}
	err = ur.Conn.Where("id = ?", data.ID).First(&record).Error
	if err != nil {
		return err
	}
	record.JumlahComment++
	err = ur.Conn.Model(&Thread{}).Where("id = ?", data.ID).Update("jumlah_comment", record.JumlahComment).Error
	return
}

func (tr *mysqlThreadRepository) DeleteThreadbyId(data threads.Core) (err error) {
	record := Thread{}
	err = tr.Conn.Where("id = ?", data.ID).Delete(&record).Error
	return
}

func (tr *mysqlThreadRepository) SearchThread(data threads.Core) (resp []threads.Core, err error) {
	record := []Thread{}
	err = tr.Conn.Select("id").Where("title LIKE ? ", "%"+data.Search+"%").Or("description LIKE ? ", "%"+data.Search+"%").Find(&record).Error
	resp = ToCoreSlice(record)
	return
}

func (tr *mysqlThreadRepository) SelectThreadAll(data threads.Core) (resp []threads.Core, err error) {
	record := []Thread{}
	err = tr.Conn.Limit(20).Offset(data.Page * 20).Find(&record).Error
	resp = ToCoreSlice(record)
	return
}

func (tr *mysqlThreadRepository) SelectThreadUser(data threads.Core) (resp []threads.Core, err error) {
	record := []Thread{}
	err = tr.Conn.Limit(20).Offset(data.Page*20).Where("user_id = ?", data.UserID).Find(&record).Error
	resp = ToCoreSlice(record)
	return
}
