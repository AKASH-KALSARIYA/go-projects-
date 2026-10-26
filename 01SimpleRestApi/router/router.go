package router

import (
	controller "github.com/AKASH-KALSARIYA/mygolang/FirstProject/handler"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api", controller.Getalltasks).Methods("GET")
	router.HandleFunc("/api/addTask", controller.CreateTask).Methods("POST")
	router.HandleFunc("/api/{id}", controller.Updateonetasks).Methods("PUT")
	router.HandleFunc("/api/{id}", controller.Deleteonetask).Methods("DELETE")

	return router
}
