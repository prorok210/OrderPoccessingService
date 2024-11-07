package main

import (
	"fmt"
	"log"

	acc "github.com/prorok210/OrderProcessingService/internal/orderAccepting"
	db "github.com/prorok210/OrderProcessingService/pkg/db"
)

func main() {
	dbCon, err := db.ConnectDB("orders.db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect to db")

	serv, err := acc.CreateAccServer(dbCon)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Create server")

	if err := serv.StartListen(); err != nil {
		log.Fatal(err)
	}
}
