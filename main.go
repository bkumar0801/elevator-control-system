package main

import (
	elv "github.com/elevator-control-system/elevator"
)

func main() {
	cs := elv.NewConstrolSystem(2, 10)
	cs.AddElevator(elv.NewElevator(100, 0))
	cs.AddElevator(elv.NewElevator(200, 0))

	cs.PickUp(3, elv.UP)
	cs.PickUp(4, elv.UP)

	for i := 0; i < 5; i++ {
		cs.Step(1)
	}
}
