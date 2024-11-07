package orderProcessing

func CreateQueue() *Queue {
	return &Queue{
		arr: make([]*Order, 1),
	}
}

type Order struct {
	Name       string
	UniqNumber int
	Status     string
}

type Queue struct {
	arr       []*Order
	lastIndex int
}

func (que *Queue) Add(ord Order) {
	if que.arr[0] == nil {
		que.arr[0] = &ord
		que.lastIndex = 0
	} else {
		que.arr = append(que.arr, &ord)
		que.lastIndex++
	}
}

func (que *Queue) Get() *Order {
	if que.arr[0] == nil {
		return nil
	} else {
		ord := que.arr[0]
		if que.lastIndex > 0 {
			que.arr = que.arr[1:]
		} else {
			que.arr[0] = nil
		}
		que.lastIndex--
		return ord
	}
}
