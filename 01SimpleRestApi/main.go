package main

import (
	"fmt"
	"log"
	"net/http"

	configs "github.com/AKASH-KALSARIYA/mygolang/FirstProject/configs"
	"github.com/AKASH-KALSARIYA/mygolang/FirstProject/router"
)

func main() {
	configs.ConnectDB()
	fmt.Println("First Project!!!")
	r := router.Router()
	fmt.Println("Server IS Getting Started...")
	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("Listenig At PORT 8000....")
}
