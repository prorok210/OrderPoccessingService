package main

import (
	"fmt"
	"log"
	"net"

	gp "github.com/prorok210/OrderProcessingService/internal/orderProcessing"
	pr "github.com/prorok210/OrderProcessingService/proto/processingProto"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := gp.CreateGRPCServer()

	pr.RegisterOrderProcessingServiceServer(s, srv)
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Сервер обработки заказов запущен на порту :50051")

	go srv.Start()

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}
