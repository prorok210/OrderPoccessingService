package orderProcessing

import (
	"fmt"
	"sync/atomic"
	"time"

	qe "github.com/prorok210/OrderProcessingService/pkg/utils"
)

func Processing(ch <-chan qe.Order, out chan<- qe.Order, ready *atomic.Bool) {
	for {
		select {
		case ord := <-ch:
			fmt.Println(ord)

			ready.Store(false)
			ord.Status = "В работе"
			someWork(&ord)
			out <- ord
			ready.Store(true)
		}
	}
}

func someWork(ord *qe.Order) {
	timer := time.NewTimer(time.Second * 20)

	<-timer.C
	ord.Status = "Готов"
}

func (serv *GrpcProcessingServer) processResults() {
	for outOrd := range serv.OutChan {
		serv.mu.Lock()
		fmt.Println(outOrd)
		serv.StatMap[outOrd.UniqNumber] = &outOrd
		serv.mu.Unlock()
	}
}

func (serv *GrpcProcessingServer) distributeOrders() {
	for {
		var order *qe.Order
		if serv.ready1.Load() || serv.ready2.Load() || serv.ready3.Load() {
			serv.mu.Lock()
			order = serv.Queue.Get()
			serv.mu.Unlock()
		} else {
			fmt.Println("Все воркеры заняты")
			time.Sleep(1 * time.Second)
			continue
		}

		if order == nil {
			time.Sleep(1000 * time.Millisecond)
			continue
		}

		fmt.Printf("Получен заказ из очереди: %+v\n", order)

		select {
		case serv.OrdChan1 <- *order:
			fmt.Printf("Заказ №%d в работе 1 вокера", order.UniqNumber)
		case serv.OrdChan2 <- *order:
			fmt.Printf("Заказ №%d в работе 2 воркера", order.UniqNumber)
		case serv.OrdChan3 <- *order:
			fmt.Printf("Заказ №%d в работе 3 воркера", order.UniqNumber)
		}
	}
}
