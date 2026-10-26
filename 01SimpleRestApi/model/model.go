package model

type Task struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}
