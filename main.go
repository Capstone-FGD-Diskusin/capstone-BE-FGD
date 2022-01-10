package main

import (
	"fmt"
	"log"

	"github.com/dragranzer/capstone-BE-FGD/config"
	"github.com/dragranzer/capstone-BE-FGD/routes"
)

func main() {
	fmt.Println("Hello world")
	config.LoadEnv()
	config.ConnectDB()
	// migrate.AutoMigrate()
	// fmt.Println(config.ENV.PORT)

	e := routes.Setup()
	log.Fatalln(e.Start(":" + config.ENV.PORT))
}
