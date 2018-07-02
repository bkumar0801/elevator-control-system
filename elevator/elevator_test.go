package elevator

import (
	"testing"
)

func TestNewElevatorInitialization(t *testing.T) {

	startingFloor := 0
	elevatorId := 1

	el := NewElevator(elevatorId, startingFloor)

	if el.CurrentFloor != startingFloor {
		t.Error("After initializiation currentFloor should be startingFloor but was:", el.CurrentFloor)
	}

	if el.Direction != IDLE {
		t.Error("After initializiation Direction should be IDLE(0) but was:", el.Direction)
	}

	if el.Goals.Len() > 0 {
		t.Error("After initializiation Goals should be empty but was:", el.Goals.Len())
	}

	if el.ID != elevatorId {
		t.Error("After initializiation Direction should be IDLE(0) but was:", el.Direction)
	}
}

func TestElevatorAddGoal(t *testing.T) {
	startingFloor := 0
	elevatorId := 1

	callingFloor := 3
	direction := UP
	priority := 5.0

	req := NewRequest(callingFloor, direction)
	el := NewElevator(elevatorId, startingFloor)

	el.AddGoal(req, priority)

	if el.Goals.Len() < 1 {
		t.Error("After AddGoal, Goals should have 1 Item but had:", el.Goals.Len())
	}

	if el.Direction != IDLE {
		t.Error("After AddGoal, elevator direction shoud be UP, but was:", el.Direction)
	}

}

func TestElevatorStep(t *testing.T) {
	startingFloor := 0
	elevatorId := 1

	callingFloor := 3
	direction := UP
	priority := 5.0

	req := NewRequest(callingFloor, direction)
	el := NewElevator(elevatorId, startingFloor)
	el.AddGoal(req, priority)

	el.Step()

	if el.Goals.Len() == 0 {
		t.Error("After Step() Goals should be still hold 1 Goal, due to not reaching goalFloor, but was:", el.Goals.Len())
	}

	if el.CurrentFloor != startingFloor+1 {
		t.Error("After Step() currentFloor should be startingFloor + 1 but was:", el.CurrentFloor)
	}

	if el.Direction != UP {
		t.Error("After initializiation Direction should be UP(1) but was:", el.Direction)
	}

}

func TestElevatorStepsUntilGoal(t *testing.T) {
	startingFloor := 0
	elevatorId := 1

	callingFloor := 3
	direction := UP
	priority := 5.0

	req := NewRequest(callingFloor, direction)
	el := NewElevator(elevatorId, startingFloor)
	el.AddGoal(req, priority)
	el.Step()
	el.Step()

	el.Step()

	if el.CurrentFloor != callingFloor {
		t.Error("After 3 x Step() currentFloor should be callingFloor(goalFloor) but was:", el.CurrentFloor)
	}

	if el.Direction != IDLE {
		t.Error("After reaching Goal, Direction should be IDLE(0) but was:", el.Direction)
	}

}

func TestElevatorStepsAfterGoal(t *testing.T) {
	startingFloor := 0
	elevatorId := 1

	callingFloor := 3
	direction := UP
	priority := 5.0

	req := NewRequest(callingFloor, direction)
	el := NewElevator(elevatorId, startingFloor)
	el.AddGoal(req, priority)
	el.Step()
	el.Step()
	el.Step()

	//Goal is reached, additional steps

	el.Step()

	if el.CurrentFloor != callingFloor {
		t.Error("After 4 x Step() currentFloor should be callingFloor(goalFloor) but was:", el.CurrentFloor)
	}

	if el.Direction != IDLE {
		t.Error("After reaching Goal, Direction should be IDLE(0) but was:", el.Direction)
	}

}

func TestElevatorDistanceToCallingFloor(t *testing.T) {
	startingFloor := 0
	elevatorId := 1

	callingFloor := 3
	direction := UP

	req := NewRequest(callingFloor, direction)
	el := NewElevator(elevatorId, startingFloor)

	distance := el.Distance(req)

	if distance != 3 {
		t.Error("Distance should be 3, but was:", distance)
	}
}

func TestElevatorPriorityOfGoals(t *testing.T) {
	startingFloor := 3
	elevatorId := 1

	reqHighPriority := NewRequest(7, UP)
	lowerPriority := NewRequest(2, DOWN)

	el := NewElevator(elevatorId, startingFloor)

	el.AddGoal(reqHighPriority, 15)
	el.AddGoal(lowerPriority, 1)

	el.Step()

	if el.CurrentFloor != 4 {
		t.Error("CurrentFloor after 1 x Step() should be 4 but was:", el.CurrentFloor)
	}

	if el.Direction != UP {
		t.Error("Direction should ne UP(1) due to higher priority, but was:", el.Direction)
	}
}

func TestElevatorDoesntMoveIfCurrentFloorIsGoalFloor(t *testing.T) {
	startingFloor := 0
	elevatorId := 1

	reqHighPriority := NewRequest(0, DOWN)

	el := NewElevator(elevatorId, startingFloor)

	el.AddGoal(reqHighPriority, 15)

	el.Step()

	if el.CurrentFloor != 0 {
		t.Error("CurrentFloor after 1 x Step() should be 0 but was:", el.CurrentFloor)
	}

	if el.Direction != IDLE {
		t.Error("Direction should ne IDLE, but was:", el.Direction)
	}
}
