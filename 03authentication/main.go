package main

import (
	"auth/config"
	"auth/router"
	"fmt"
)

func main() {
	fmt.Println("AUTH PROJECT STARTED")

	config.ConnectDB()

	r := router.Router()
	r.Run(":4000")
}
