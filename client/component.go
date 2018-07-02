package client

/* FB ... EB */
const (
	FB = 0 //FloorButton
	EB = 1 //ElevatorButton
)

/*
Component ...
*/
type Component interface {
	RequestPickup() Request
}

/*
Request ...
*/
type Request struct {
	PickupFloor int
	Direction   int
	FromDevice  Device
}

/*
Device ...
*/
type Device struct {
	Type int
	ID   int
}

/*
FloorButton ...
*/
type FloorButton struct {
	Pickup Request
}

/*
ElevatorButton ...
*/
type ElevatorButton struct {
	Pickup Request
}

/*
RequestPickup ...
*/
func (fb *FloorButton) RequestPickup() Request {
	return fb.Pickup
}

/*
RequestPickup ...
*/
func (cb *ElevatorButton) RequestPickup() Request {
	return cb.Pickup
}
