package migrate

import (
	"github.com/dragranzer/capstone-BE-FGD/config"
	_comment_data "github.com/dragranzer/capstone-BE-FGD/features/comments/data"
	_like_data "github.com/dragranzer/capstone-BE-FGD/features/likes/data"
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

	config.DB.AutoMigrate(&_user_data.User{}, &_thread_data.Thread{}, &_comment_data.Comment{}, &_like_data.Like{})

	pass1, _ := HashPassword("pass1")
	user1 := _user_data.User{
		Username: "Zehan",
		Email:    "zehan@gmail.com",
		Password: pass1,
	}

	pass2, _ := HashPassword("pass2")
	user2 := _user_data.User{
		Username: "Ivan",
		Email:    "ivan@gmail.com",
		Password: pass2,
	}

	pass3, _ := HashPassword("pass3")
	user3 := _user_data.User{
		Username: "Faris",
		Email:    "faris@gmail.com",
		Password: pass3,
	}

	thread1 := _thread_data.Thread{
		Title:       "Mengapa GO lebih baik?",
		Description: "Go merupakan bahasa yang dikembangkan oleh google, oleh karena itu...",
		UserID:      1,
		Like:        2,
	}

	thread2 := _thread_data.Thread{
		Title:       "Bahas tuntas Meta blacklist",
		Description: "Walaupun heronya cuman itu itu aja tetapi...",
		UserID:      1,
		Like:        2,
	}

	comment1 := _comment_data.Comment{
		Comment:  "GO emang didesain seperti itu gan",
		UserID:   2,
		ThreadID: 1,
	}

	comment2 := _comment_data.Comment{
		Comment:   "Jawaban yang tidak menjawab",
		UserID:    3,
		ThreadID:  1,
		CommentID: 1,
	}

	comment3 := _comment_data.Comment{
		Comment:  "Viva RRQ",
		UserID:   1,
		ThreadID: 2,
	}

	like1 := _comment_data.Comment{
		UserID:   1,
		ThreadID: 2,
	}

	like2 := _comment_data.Comment{
		UserID:   1,
		ThreadID: 1,
	}

	like3 := _comment_data.Comment{
		UserID:   2,
		ThreadID: 1,
	}

	like4 := _comment_data.Comment{
		UserID:   2,
		ThreadID: 2,
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

	if err := config.DB.Create(&thread1).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&thread2).Error; err != nil {
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
}
