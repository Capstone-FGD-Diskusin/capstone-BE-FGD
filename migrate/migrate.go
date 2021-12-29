package migrate

import (
	"github.com/dragranzer/capstone-BE-FGD/config"
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

	config.DB.AutoMigrate(&_user_data.User{})

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

	if err := config.DB.Create(&user1).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&user2).Error; err != nil {
		panic(err)
	}

	if err := config.DB.Create(&user3).Error; err != nil {
		panic(err)
	}
}
