package orderProcessing

import (
	"sync"
	"sync/atomic"

	qe "github.com/prorok210/OrderProcessingService/pkg/utils"
	pt "github.com/prorok210/OrderProcessingService/proto/processingProto"
)

type GrpcProcessingServer struct {
	pt.UnimplementedOrderProcessingServiceServer
	Queue     *qe.Queue
	StatMap   map[int]*qe.Order
	mu        *sync.Mutex
	OrdChan1  chan qe.Order
	OrdChan2  chan qe.Order
	OrdChan3  chan qe.Order
	ready1    *atomic.Bool
	ready2    *atomic.Bool
	ready3    *atomic.Bool
	OutChan   chan qe.Order
	FinishJob chan struct{}
}

func CreateGRPCServer() *GrpcProcessingServer {
	serv := &GrpcProcessingServer{
		Queue:     qe.CreateQueue(),
		StatMap:   make(map[int]*qe.Order),
		mu:        &sync.Mutex{},
		OrdChan1:  make(chan qe.Order),
		OrdChan2:  make(chan qe.Order),
		OrdChan3:  make(chan qe.Order),
		ready1:    &atomic.Bool{},
		ready2:    &atomic.Bool{},
		ready3:    &atomic.Bool{},
		OutChan:   make(chan qe.Order),
		FinishJob: make(chan struct{}),
	}
	serv.ready1.Store(true)
	serv.ready2.Store(true)
	serv.ready3.Store(true)
	return serv
}

func (serv *GrpcProcessingServer) Start() {
	go Processing(serv.OrdChan1, serv.OutChan, serv.ready1)
	go Processing(serv.OrdChan2, serv.OutChan, serv.ready2)
	go Processing(serv.OrdChan3, serv.OutChan, serv.ready3)

	go serv.processResults()

	go serv.distributeOrders()
}
