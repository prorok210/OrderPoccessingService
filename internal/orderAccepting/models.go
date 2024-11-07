package orderAccepting

type Order struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

const (
	Accept        string = "Принят"
	InProccessing string = "Взят в работу"
	Ready         string = "Готов"
	Canceled      string = "Отменен"
)
