package main

import (
	"fmt"
	"log"
	"net/http"

	//"load-management/models"
	"github.com/md96/load-management/db"
	"github.com/md96/load-management/restcontroller"
	"github.com/md96/load-management/websocket"
)

func main() {

	//DB Connection Initialization
	db.DBInit()

	fmt.Println("Load management System micro service ")
	http.HandleFunc("/stationcreation", restcontroller.Stationcreation)
	http.HandleFunc("/getallstations", restcontroller.Getallstations)
	http.HandleFunc("/deletestation", restcontroller.Deletestation)
	fmt.Println("Load Management Service is started on 8082 Port")
	http.HandleFunc("/ws/", websocket.Websockethandler)

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
