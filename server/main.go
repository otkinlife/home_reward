package main

import (
	"fmt"
	"home-reward/server/api"
	"home-reward/server/base"
	"home-reward/server/config"
	"net/http"
	"os"
)

func main() {
	config.InitConfig()
	base.InitMySQL()

	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", api.GetTaskList)
	mux.HandleFunc("/task/create", api.CreateTask)
	mux.HandleFunc("/task/del", api.DelTask)
	mux.HandleFunc("/task/get", api.GetTask)
	mux.HandleFunc("/task/cancel_get", api.CancelGetTask)
	mux.HandleFunc("/task/finish", api.FinishTask)
	mux.HandleFunc("/task/cancel_finish", api.CancelFinishTask)

	mux.HandleFunc("/wishes", api.GetWishList)
	mux.HandleFunc("/wish/create", api.CreateWish)
	mux.HandleFunc("/wish/del", api.DelWish)
	mux.HandleFunc("/wish/finish", api.FinishWish)
	mux.HandleFunc("/wish/cancel_finish", api.CancelFinishWish)

	mux.HandleFunc("/character", api.GetCharacter)
	mux.HandleFunc("/bind", api.BindCharacterWithIP)
	mux.HandleFunc("/unbind", api.UnBindCharacterWithIP)

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("../client"))))
	fmt.Println("server start at port 8999")
	err := http.ListenAndServe(":8999", mux)
	if err != nil {
		os.Exit(1)
	}
}
