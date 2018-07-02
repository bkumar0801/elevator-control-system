package elevator

import (
	"math"
)

/* FB ... EB */
const (
	FB = 0 //FloorButton
	EB = 1 //ElevatorButton
)

/*
ControlSystemInterface ...
*/
type ControlSystemInterface interface {
	Status() *ControlSystemStatus
	Update(id int, floor int, floorToVisit int) bool
	Pickup(floor, direction int)
	Step()
}

/*
ControlSystem ...
*/
type ControlSystem struct {
	Elevators    []*Elevator
	MaxElevators int
	MaxFloors    int
}

/*
ControlSystemStatus ...
*/
type ControlSystemStatus []*Elevator

/*
Len ...
*/
func (css *ControlSystemStatus) Len() int {
	return len(*css)
}

/*
GetStatusAtIndex ...
*/
func (css *ControlSystemStatus) GetStatusAtIndex(index int) *Elevator {
	cpy := *css
	return cpy[index]
}

/*
NewConstrolSystem ...
*/
func NewConstrolSystem(maxElevators, maxFloors int) *ControlSystem {
	return &ControlSystem{
		Elevators:    []*Elevator{},
		MaxElevators: maxElevators,
		MaxFloors:    maxFloors,
	}
}

func calculateSuitabilityScore(numberOfFloors int, elevator *Elevator, req *Request) float64 {

	if goal := elevator.Goals.Find(req.Floor); goal != nil {
		return goal.Priority
	}

	direction := elevator.Direction
	distance := elevator.Distance(req)

	if direction*distance >= 0 {
		if direction == req.Direction {
			return float64(numberOfFloors+2) - math.Abs(float64(distance))
		}
		return float64(numberOfFloors+1) - math.Abs(float64(distance))
	}
	return 1.0
}

/*
AddElevator ...
*/
func (cs *ControlSystem) AddElevator(elevator *Elevator) {
	cs.Elevators = append(cs.Elevators, elevator)
}

/*
PickUp ...
*/
func (cs *ControlSystem) PickUp(floor, direction, device, id int) {
	req := NewRequest(floor, direction)

	highestScore := 0.0
	var chosenElevator *Elevator

	for _, elevator := range cs.Elevators {
		score := calculateSuitabilityScore(cs.MaxFloors, elevator, req)

		if score > highestScore {
			highestScore = score
			chosenElevator = elevator
		}
		if device == EB && id == elevator.ID {
			chosenElevator = elevator
		}
	}
	chosenElevator.AddGoal(req, highestScore)
}

/*
Status ...
*/
func (cs *ControlSystem) Status() *ControlSystemStatus {
	css := ControlSystemStatus(cs.Elevators)
	return &css
}

/*
Step ...
*/
func (cs *ControlSystem) Step(steps int) {
	for i := 0; i < steps; i++ {
		for _, elevator := range cs.Elevators {
			elevator.Step()
		}
	}
}
