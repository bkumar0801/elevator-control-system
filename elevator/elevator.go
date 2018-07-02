package elevator

import (
	"fmt"

	"github.com/elevator-control-system/queue"
)

/* DOWN ... IDLE ... UP */
const (
	DOWN = -1
	IDLE = 0
	UP   = 1
)

/*
Elevator ...
*/
type Elevator struct {
	ID           int
	CurrentFloor int
	Goals        *queue.PriorityQueue
	Direction    int
}

/*
NewElevator ...
*/
func NewElevator(id, startingFloor int) *Elevator {
	return &Elevator{
		ID:           id,
		CurrentFloor: startingFloor,
		Goals:        queue.NewPriorityQueue(),
		Direction:    IDLE,
	}
}

/*
Step ...
*/
func (e *Elevator) Step() {

	if e.Goals.Len() == 0 {
		e.Direction = IDLE
		return
	}

	currentGoal := e.Goals.Peek()
	e.move(currentGoal.Floor)
	fmt.Printf("Elevator ID:%d \n\t\t Current Floor: %d \n\t\t Direction: %d \n\t\t Goals: %v\n", e.ID, e.CurrentFloor, e.Direction, e.Goals.PrintGoals())
}

/*
Distance ...
*/
func (e *Elevator) Distance(req *Request) int {
	return req.Floor - e.CurrentFloor
}

/*
AddGoal ...
*/
func (e *Elevator) AddGoal(req *Request, priority float64) {

	if req.Floor == e.CurrentFloor {
		return
	}

	goal := queue.NewGoal(req.Floor, priority)
	e.Goals.PushGoal(goal)
}

func (e *Elevator) move(goalFloor int) {

	if goalFloor > e.CurrentFloor {
		e.Direction = UP
		e.CurrentFloor++

	} else {
		e.Direction = DOWN
		e.CurrentFloor--
	}

	if goalFloor == e.CurrentFloor {
		e.Goals.Pop()
		if e.Goals.Len() == 0 {
			e.Direction = IDLE
		}
	}
}

/*
String ...
*/
func (e *Elevator) String() string {
	return fmt.Sprintf("Elevator ID: {%v} on Floor: {%v} in direction: {%v} going to floor:%v  Goals: %v\n", e.ID, e.CurrentFloor, e.Direction, e.Goals.Peek(), e.Goals.PrintGoals())
}
