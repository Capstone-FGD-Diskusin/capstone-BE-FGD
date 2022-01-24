package migrate

import (
	"github.com/dragranzer/capstone-BE-FGD/config"
	_admin_data "github.com/dragranzer/capstone-BE-FGD/features/admins/data"
	_category_data "github.com/dragranzer/capstone-BE-FGD/features/categories/data"
	_comment_data "github.com/dragranzer/capstone-BE-FGD/features/comments/data"
	_detail_thread_data "github.com/dragranzer/capstone-BE-FGD/features/detail_threads/data"
	_favorite_data "github.com/dragranzer/capstone-BE-FGD/features/favorites/data"
	_follower_data "github.com/dragranzer/capstone-BE-FGD/features/followers/data"
	_like_data "github.com/dragranzer/capstone-BE-FGD/features/likes/data"
	_message_data "github.com/dragranzer/capstone-BE-FGD/features/messages/data"
	_tag_data "github.com/dragranzer/capstone-BE-FGD/features/tags/data"
	_thread_data "github.com/dragranzer/capstone-BE-FGD/features/threads/data"
	_user_data "github.com/dragranzer/capstone-BE-FGD/features/users/data"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func AutoMigrate() {
	if err := config.DB.Exec("DROP TABLE IF EXISTS users").Error; err != nil {
		panic(err)
	}

	if err := config.DB.Exec("DROP TABLE IF EXISTS threads").Error; err != nil {
		panic(err)
	}

	if err := config.DB.Exec("DROP TABLE IF EXISTS comments").Error; err != nil {
		panic(err)
	}

	if err := config.DB.Exec("DROP TABLE IF EXISTS likes").Error; err != nil {
		panic(err)
	}

	if err := config.DB.Exec("DROP TABLE IF EXISTS favorites").Error; err != nil {
		panic(err)
	}

	if err := config.DB.Exec("DROP TABLE IF EXISTS tags").Error; err != nil {
		panic(err)
	}

	if err := config.DB.Exec("DROP TABLE IF EXISTS categories").Error; err != nil {
		panic(err)
	}

	if err := config.DB.Exec("DROP TABLE IF EXISTS followers").Error; err != nil {
		panic(err)
	}

	if err := config.DB.Exec("DROP TABLE IF EXISTS admins").Error; err != nil {
		panic(err)
	}

	if err := config.DB.Exec("DROP TABLE IF EXISTS messages").Error; err != nil {
		panic(err)
	}

	config.DB.AutoMigrate(&_user_data.User{}, &_thread_data.Thread{}, &_comment_data.Comment{}, &_like_data.Like{},
		&_favorite_data.Favorite{}, &_detail_thread_data.Detail_thread{}, &_follower_data.Follower{}, &_admin_data.Admin{},
		&_tag_data.Tag{}, &_category_data.Category{}, &_message_data.Message{})

	pass1, _ := HashPassword("pass1")
	user1 := _user_data.User{
		Username:  "Zehan",
		Email:     "zehan@gmail.com",
		Password:  pass1,
		Role:      "user",
		Following: 2,
		Follower:  2,
		Alamat:    "Jawa Timur",
		Gender:    "L",
		Phone:     "081234",
	}

	pass2, _ := HashPassword("pass2")
	user2 := _user_data.User{
		Username:  "Ivan",
		Email:     "ivan@gmail.com",
		Password:  pass2,
		Role:      "user",
		Following: 1,
		Follower:  1,
		Alamat:    "Jawa Timur",
		Gender:    "L",
		Phone:     "081234",
	}

	pass3, _ := HashPassword("pass3")
	user3 := _user_data.User{
		Username:  "Faris",
		Email:     "faris@gmail.com",
		Password:  pass3,
		Role:      "user",
		Following: 1,
		Follower:  1,
		Alamat:    "Jawa Timur",
		Gender:    "L",
		Phone:     "081234",
	}

	pass4, _ := HashPassword("pass4")
	user4 := _user_data.User{
		Username:   "Moderator1",
		Email:      "mod1@gmail.com",
		Password:   pass4,
		Role:       "moderator",
		CategoryID: 1,
	}

	pass5, _ := HashPassword("pass5")
	user5 := _user_data.User{
		Username: "Admin1",
		Email:    "admin1@gmail.com",
		Password: pass5,
		Role:     "admin",
	}

	thread1 := _thread_data.Thread{
		Title:       "Mengapa GO lebih baik?",
		Description: "Go merupakan bahasa yang dikembangkan oleh google, oleh karena itu...",
		UserID:      2,
		Like:        2,
		CategoryID:  4,
	}

	thread2 := _thread_data.Thread{
		Title:       "Bahas tuntas Meta blacklist",
		Description: "Walaupun heronya cuman itu itu aja tetapi...",
		UserID:      3,
		Like:        2,
		CategoryID:  4,
	}

	thread3 := _thread_data.Thread{
		Title:       "Google",
		Description: "Google merupakan perusahaan raksasa yang bergerak pada bidang IT",
		UserID:      2,
		Like:        2,
		CategoryID:  3,
	}

	comment1 := _comment_data.Comment{
		Comment:  "GO emang didesain seperti itu gan",
		UserID:   2,
		ThreadID: 1,
	}

	comment2 := _comment_data.Comment{
		Comment:   "Balasan dari 'Go emang didesain...'",
		UserID:    3,
		ThreadID:  1,
		CommentID: 1,
	}

	comment3 := _comment_data.Comment{
		Comment:  "Viva RRQ",
		UserID:   1,
		ThreadID: 2,
	}

	comment4 := _comment_data.Comment{
		Comment:   "Balasan dari 'Go emang didesain... V2'",
		UserID:    3,
		ThreadID:  1,
		CommentID: 1,
	}

	message1 := _message_data.Message{
		Text:          "Message 1 merupakan blablabla",
		CategoryName:  "Computer",
		ThreadID:      1,
		CommentID:     0,
		AdminID:       5,
		ModeratorName: "Moderator1",
	}

	like1 := _like_data.Like{
		UserID:   1,
		ThreadID: 2,
	}

	like2 := _like_data.Like{
		UserID:   1,
		ThreadID: 1,
	}

	like3 := _like_data.Like{
		UserID:   2,
		ThreadID: 1,
	}

	like4 := _like_data.Like{
		UserID:   2,
		ThreadID: 2,
	}

	favorite1 := _favorite_data.Favorite{
		UserID:   2,
		ThreadID: 2,
	}

	favorite2 := _favorite_data.Favorite{
		UserID:   1,
		ThreadID: 2,
	}

	tag1 := _tag_data.Tag{
		Name: "programming",
	}

	tag2 := _tag_data.Tag{
		Name: "computer-science",
	}

	tag3 := _tag_data.Tag{
		Name: "MLBB",
	}

	detail_thread1 := _detail_thread_data.Detail_thread{
		TagID:    1,
		ThreadID: 1,
	}

	detail_thread2 := _detail_thread_data.Detail_thread{
		TagID:    2,
		ThreadID: 1,
	}

	detail_thread3 := _detail_thread_data.Detail_thread{
		TagID:    3,
		ThreadID: 2,
	}

	category1 := _category_data.Category{
		Name: "Politik",
	}

	category2 := _category_data.Category{
		Name: "Hiburan",
	}

	category3 := _category_data.Category{
		Name: "Computer",
	}

	category4 := _category_data.Category{
		Name: "Game",
	}

	follower1 := _follower_data.Follower{
		FollowingID: 1,
		FollowedID:  2,
	}

	follower2 := _follower_data.Follower{
		FollowingID: 1,
		FollowedID:  3,
	}

	follower3 := _follower_data.Follower{
		FollowingID: 2,
		FollowedID:  1,
	}

	follower4 := _follower_data.Follower{
		FollowingID: 3,
		FollowedID:  1,
	}

	if err := config.DB.Create(&user1).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&user2).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&user3).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&user4).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&user5).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&thread1).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&thread2).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&thread3).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&comment1).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&comment2).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&comment3).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&comment4).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&like1).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&like2).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&like3).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&like4).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&favorite1).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&favorite2).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&tag1).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&tag2).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&tag3).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&detail_thread1).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&detail_thread2).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&detail_thread3).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&category1).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&category2).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&category3).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&category4).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&message1).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&follower1).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&follower2).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&follower3).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&follower4).Error; err != nil {
		panic(err)
	}
}
