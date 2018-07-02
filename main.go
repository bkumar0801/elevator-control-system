package main

import (
	elv "github.com/elevator-control-system/elevator"
)

func main() {
	c := elv.NewConstrolSystem(2, 10)
	c.AddElevator(elv.NewElevator(100, 0))
	c.AddElevator(elv.NewElevator(200, 0))

	c.PickUp(3, elv.UP)
	c.PickUp(4, elv.UP)

	for i := 0; i < 5; i++ {
		c.Step(1)
	}
}
