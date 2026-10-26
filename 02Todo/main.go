package main

import (
	"todo/connfigs"
	router "todo/router"
)

func main() {
	connfigs.ConnectDB()
	r := router.Router()
	r.Run(":4000")
}
