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
	mux.HandleFunc("/create_task", api.CreateTask)
	mux.HandleFunc("/del_task", api.DelTask)
	mux.HandleFunc("/do_task", api.GetTask)
	mux.HandleFunc("/undo_task", api.CancelGetTask)
	mux.HandleFunc("/done_task", api.FinishTask)

	mux.HandleFunc("/wishes", api.GetWishList)
	mux.HandleFunc("/create_wish", api.CreateWish)
	mux.HandleFunc("/del_wish", api.DelWish)
	mux.HandleFunc("/buy_wish", api.FinishWish)
	mux.HandleFunc("/cancel_buy_wish", api.CancelFinishWish)

	mux.HandleFunc("/character", api.GetCharacter)

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("../client"))))
	fmt.Println("server start at port 8999")
	err := http.ListenAndServe(":8999", mux)
	if err != nil {
		os.Exit(1)
	}
}
