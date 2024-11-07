package orderAccepting

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	db "github.com/prorok210/OrderProcessingService/pkg/db"
	pt "github.com/prorok210/OrderProcessingService/proto/processingProto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderHanlder struct {
	db *sql.DB
}

func (h *OrderHanlder) CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var ord Order
	err := json.NewDecoder(r.Body).Decode(&ord)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if ord.Name == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	id, err := db.InsertValue(h.db, ord.Name, Accept)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		http.Error(w, "Error connecting", http.StatusInternalServerError)
		log.Fatalf("Error connecting: %v", err)
	}
	defer conn.Close()

	client := pt.NewOrderProcessingServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pt.OrderInfo{
		UniqNumber: int32(id),
		OrderName:  ord.Name,
	}

	_, err = client.AddOrder(ctx, req)
	if err != nil {
		http.Error(w, "Error when calling ProcessOrder", http.StatusInternalServerError)
		log.Fatalf("Error when calling ProcessOrder: %v", err)
	}

}
