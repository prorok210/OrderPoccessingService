package orderProcessing

import (
	"context"
	"errors"
	"fmt"

	qe "github.com/prorok210/OrderProcessingService/pkg/utils"
	pt "github.com/prorok210/OrderProcessingService/proto/processingProto"
)

func (serv *GrpcProcessingServer) AddOrder(ctx context.Context, req *pt.OrderInfo) (*pt.OrderStatus, error) {
	if req == nil {
		return nil, errors.New("Nil request")
	}

	serv.mu.Lock()
	defer serv.mu.Unlock()

	if serv.Queue == nil {
		serv.Queue = qe.CreateQueue()
	}

	ord := qe.Order{
		Name:       req.GetOrderName(),
		UniqNumber: int(req.GetUniqNumber()),
		Status:     "Принят",
	}

	if serv.StatMap[ord.UniqNumber] != nil {
		fmt.Println("Заказ с таким номер уже есть в системе")
		return &pt.OrderStatus{
			UniqNumber: int32(ord.UniqNumber),
			Status:     "Уже в обработке",
		}, nil
	}

	serv.Queue.Add(ord)
	serv.StatMap[ord.UniqNumber] = &ord

	return &pt.OrderStatus{
		UniqNumber: int32(ord.UniqNumber),
		Status:     ord.Status,
	}, nil
}

func (serv *GrpcProcessingServer) CheckStatus(ctx context.Context, req *pt.OrderInfo) (*pt.OrderStatus, error) {
	ord := serv.StatMap[int(req.GetUniqNumber())]
	if ord == nil {
		return &pt.OrderStatus{
			Status: "Заказа с таким номером нет в системе",
		}, nil
	}
	return &pt.OrderStatus{
		UniqNumber: int32(ord.UniqNumber),
		Status:     ord.Status,
	}, nil
}
