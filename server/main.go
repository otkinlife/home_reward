package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"home-reward/server/api"
	"home-reward/server/dao"
	"home-reward/server/object"
	"net/http"
	"os"
)

const dataPath = "../data"

func main() {
	dao.InitMySql()
	object.InitConfig()
	object.InitTaskList()
	object.InitProductList()
	object.InitCharacters()

	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", api.GetTaskList)
	mux.HandleFunc("/create_task", api.CreateTask)
	mux.HandleFunc("/del_task", api.DelTask)
	mux.HandleFunc("/do_task", api.DoTask)
	mux.HandleFunc("/undo_task", api.UnDoTask)
	mux.HandleFunc("/done_task", api.DoneTask)

	mux.HandleFunc("/products", api.GetProductList)
	mux.HandleFunc("/create_product", api.CreateProduct)
	mux.HandleFunc("/del_product", api.DelProduct)
	mux.HandleFunc("/buy_product", api.BuyProduct)
	mux.HandleFunc("/cancel_buy_product", api.CancelBuyProduct)

	mux.HandleFunc("/character", api.GetCharacter)

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("../client"))))
	fmt.Println("server start at port 8999")
	err := http.ListenAndServe(":8999", mux)
	if err != nil {
		os.Exit(1)
	}
}
