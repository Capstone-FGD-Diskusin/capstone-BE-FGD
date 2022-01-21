package data

import (
	"github.com/dragranzer/capstone-BE-FGD/features/followers"
	"gorm.io/gorm"
)

type mysqlFollowerRepository struct {
	Conn *gorm.DB
}

func NewFollowerRepository(conn *gorm.DB) followers.Data {
	return &mysqlFollowerRepository{
		Conn: conn,
	}
}

func (fr *mysqlFollowerRepository) InsertFollow(data followers.Core) (err error) {
	recordData := fromCore(data)
	err = fr.Conn.Create(&recordData).Error
	if err != nil {
		return err
	}
	return err
}

func (fr *mysqlFollowerRepository) DeleteFollow(data followers.Core) (err error) {
	recordData := fromCore(data)
	err = fr.Conn.Where("following_id = ?", data.FollowingID).Where("followed_id = ?", data.FollowedID).Delete(&recordData).Error
	if err != nil {
		return err
	}
	return err
}

func (fr *mysqlFollowerRepository) SelectFollowing(data followers.Core) (resp []followers.Core, err error) {
	recordData := fromCore(data)
	// fmt.Println(recordData)
	listFollowing := []Follower{}
	err = fr.Conn.Where("following_id = ?", recordData.FollowingID).Find(&listFollowing).Error
	// fmt.Println(listFollowing)
	resp = toCoreList(listFollowing)
	return resp, err
}

func (fr *mysqlFollowerRepository) SelectFollowed(data followers.Core) (resp []followers.Core, err error) {
	recordData := fromCore(data)
	// fmt.Println(recordData)
	listFollowed := []Follower{}
	err = fr.Conn.Where("followed_id = ?", recordData.FollowedID).Find(&listFollowed).Error
	// fmt.Println(listFollowed)
	resp = toCoreList(listFollowed)
	return resp, err
}
