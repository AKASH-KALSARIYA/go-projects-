package controller

import (
	"encoding/json"

	"net/http"

	"github.com/AKASH-KALSARIYA/mygolang/FirstProject/model"
	"github.com/AKASH-KALSARIYA/mygolang/FirstProject/repository"
	"github.com/gorilla/mux"
)

// --------------------------------------------------------------------------------------------
// Actual functions -controll a function like create/update/delete/read
func Getalltasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	alltasks := repository.Getalltask()
	json.NewEncoder(w).Encode(alltasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task model.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		panic(err)
	}

	repository.Insterttask(task)

	json.NewEncoder(w).Encode(task)
}

func Updateonetasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var task model.Task
	json.NewDecoder(r.Body).Decode(&task)
	updated := repository.Updatetask(params["id"], task)

	json.NewEncoder(w).Encode(updated)
}

func Deleteonetask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	repository.Deletetask(params["id"])

	json.NewEncoder(w).Encode(params["id"])

}
