package elevator

import (
	"testing"
)

func TestNewControlSystem(t *testing.T) {

	maxElevators := 1
	maxFloors := 3
	expectedElevator := NewElevator(0, 0)
	cs := NewConstrolSystem(maxElevators, maxFloors)
	cs.AddElevator(NewElevator(0, 0))
	csStatus := ControlSystemStatus([]*Elevator{expectedElevator})
	expectedStatus := &csStatus

	stateAfterCreation := cs.Status()

	if stateAfterCreation == nil {
		t.Error("After creation Status must not be nil")
	}

	if stateAfterCreation.Len() != cs.MaxElevators {
		t.Errorf("Number of status for elevator should be number of elevators, but was:%v", stateAfterCreation.Len())
	}

	areEqual := controlSystemStatusAreEqual(*stateAfterCreation, *expectedStatus)
	if !areEqual {
		t.Errorf("stateAfterCreation must be equal to expectedStatus, but was:%v", areEqual)
	}
}

func TestControlSystem(t *testing.T) {

	maxElevators := 1
	maxFloors := 3

	expectedElevator := NewElevator(0, 0)
	req := NewRequest(2, UP)

	expectedElevator.AddGoal(req, 1.7)
	expectedElevator.Step()

	cs := NewConstrolSystem(maxElevators, maxFloors)
	cs.AddElevator(NewElevator(0, 0))
	cs.PickUp(2, UP)
	cs.Step(1)

	csStatus := ControlSystemStatus([]*Elevator{expectedElevator})
	expectedStatus := &csStatus

	stateAfterCreation := cs.Status()

	if stateAfterCreation == nil {
		t.Error("After creation Status must not be nil")
	}

	if stateAfterCreation.Len() != cs.MaxElevators {
		t.Errorf("Number of status for elevator should be number of elevators, but was:%v", stateAfterCreation.Len())
	}

	areEqual := controlSystemStatusAreEqual(*stateAfterCreation, *expectedStatus)
	if !areEqual {
		t.Errorf("stateAfterCreation must be equal to expectedStatus, but was:%v", areEqual)
	}
}

func TestRequestScoring(t *testing.T) {
	maxElevators := 2
	numOfFloors := 5

	cs := NewConstrolSystem(maxElevators, numOfFloors)
	cs.AddElevator(NewElevator(0, 0))
	cs.AddElevator(NewElevator(1, 0))
	cs.PickUp(5, DOWN)
	cs.Step(5)

	cs.PickUp(4, DOWN)
	cs.Step(5)

	elevators := cs.Status()
	elvFirst := elevators.GetStatusAtIndex(0)
	elvSecond := elevators.GetStatusAtIndex(elevators.Len() - 1)

	if elvFirst.CurrentFloor != 4 && elvSecond.CurrentFloor != 0 {
		t.Errorf("First elevator should be on floor 4 due to scoring and second elevator should be idle on starting floor")
	}

}

func controlSystemStatusAreEqual(a, b ControlSystemStatus) bool {
	if a.Len() != b.Len() {
		return false
	}

	for i, v := range a {
		if b[i].ID != v.ID && b[i].Direction != v.Direction &&
			b[i].CurrentFloor != v.CurrentFloor && b[i].Goals.Len() != v.Goals.Len() { //Due to time restrictions only testing on len of Goals
			return false
		}
	}
	return true
}
