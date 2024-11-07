package orderAccepting

import (
	"database/sql"
	"errors"
	"net"
	"net/http"
)

type AcceptingServer struct {
	conn *net.Listener
	db   *sql.DB
}

func CreateAccServer(db *sql.DB) (*AcceptingServer, error) {
	conn, err := net.Listen("tcp", ":8080")
	if err != nil {
		return nil, err
	}

	return &AcceptingServer{
		conn: &conn,
		db:   db,
	}, nil
}

func (s *AcceptingServer) StartListen() error {
	if s.conn == nil {
		return errors.New("Conn is nil")
	}

	mux := http.NewServeMux()

	orderHandler := OrderHanlder{
		db: s.db,
	}

	mux.HandleFunc("/create-order", orderHandler.CreateOrder)

	return http.Serve(*s.conn, mux)
}
