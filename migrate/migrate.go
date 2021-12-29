package migrate

import (
	"github.com/dragranzer/capstone-BE-FGD/config"
	_user_data "github.com/dragranzer/capstone-BE-FGD/features/users/data"
)

func AutoMigrate() {
	if err := config.DB.Exec("DROP TABLE IF EXISTS users").Error; err != nil {
		panic(err)
	}

	config.DB.AutoMigrate(&_user_data.User{})
}
