package elevator

type Request struct {
	Floor     int
	Direction int
}

func NewRequest(currentFloor, direction int) *Request {
	return &Request{
		Floor:     currentFloor,
		Direction: direction,
	}
}
