package data

import (
	"github.com/dragranzer/capstone-BE-FGD/features/comments"
	"gorm.io/gorm"
)

type mysqlCommentRepository struct {
	Conn *gorm.DB
}

func NewCommentRepository(conn *gorm.DB) comments.Data {
	return &mysqlCommentRepository{
		Conn: conn,
	}
}

func (cr *mysqlCommentRepository) InsertComment(data comments.Core) (err error) {
	record := FromCore(data)
	err = cr.Conn.Create(&record).Error
	return
}

func (cr *mysqlCommentRepository) SelectCommentsThread(data comments.Core) (resp []comments.Core, err error) {
	record := []Comment{}
	err = cr.Conn.Limit(20).Offset(data.Page*20).Where("thread_id = ? AND comment_id = ?", data.ThreadID, 0).Find(&record).Error
	if err != nil {
		return resp, err
	}
	for _, value := range record {
		resp = append(resp, ToCore(value))
	}
	return
}

func (cr *mysqlCommentRepository) DeleteCommentbyId(data comments.Core) (err error) {
	record := Comment{}
	err = cr.Conn.Where("id = ?", data.ID).Delete(&record).Error
	return
}

func (cr *mysqlCommentRepository) SelectCommentbyId(data comments.Core) (resp comments.Core, err error) {
	record := Comment{}
	err = cr.Conn.Where("id = ?", data.ID).Find(&record).Error
	if err != nil {
		return resp, err
	}
	resp = ToCore(record)
	return resp, err
}

func (cr *mysqlCommentRepository) DeleteCommentbyThreadId(data comments.Core) (err error) {
	record := Comment{}
	err = cr.Conn.Where("thread_id = ?", data.ThreadID).Delete(&record).Error
	return
}

func (cr *mysqlCommentRepository) SelectBalasanCommentbyId(data comments.Core) (resp []comments.Core, err error) {
	record := []Comment{}
	err = cr.Conn.Where("comment_id = ?", data.CommentID).Find(&record).Error
	if err != nil {
		return resp, err
	}
	for _, value := range record {
		resp = append(resp, ToCore(value))
	}
	return
}

func (tr *mysqlCommentRepository) SearchThreadbyComment(data comments.Core) (resp []comments.Core, err error) {
	record := []Comment{}
	err = tr.Conn.Select("thread_id").Where("comment LIKE ? ", "%"+data.Search+"%").Find(&record).Error
	resp = ToCoreSlice(record)
	return
}
