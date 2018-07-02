package main

import (
	"time"

	"github.com/elevator-control-system/client"
	elv "github.com/elevator-control-system/elevator"
)

func main() {
	cs := elv.NewConstrolSystem(2, 10)
	cs.AddElevator(elv.NewElevator(100, 0))
	cs.AddElevator(elv.NewElevator(200, 0))

	floorButton1 := &client.FloorButton{
		Pickup: client.Request{
			PickupFloor: 3,
			Direction:   elv.UP,
			FromDevice: client.Device{
				Type: client.FB,
				ID:   3,
			},
		},
	}

	floorButton2 := &client.FloorButton{
		Pickup: client.Request{
			PickupFloor: 4,
			Direction:   elv.UP,
			FromDevice: client.Device{
				Type: client.FB,
				ID:   4,
			},
		},
	}

	floorButton3 := &client.FloorButton{
		Pickup: client.Request{
			PickupFloor: 5,
			Direction:   elv.UP,
			FromDevice: client.Device{
				Type: client.FB,
				ID:   5,
			},
		},
	}

	floorButton4 := &client.FloorButton{
		Pickup: client.Request{
			PickupFloor: 2,
			Direction:   elv.DOWN,
			FromDevice: client.Device{
				Type: client.FB,
				ID:   2,
			},
		},
	}

	elevatorButton1 := &client.ElevatorButton{
		Pickup: client.Request{
			PickupFloor: 2,
			Direction:   elv.UP,
			FromDevice: client.Device{
				Type: client.EB,
				ID:   200,
			},
		},
	}

	mergedRequests := client.Merge(
		client.Subscribe(floorButton1),
		client.Subscribe(floorButton2),
		client.Subscribe(floorButton3),
		client.Subscribe(floorButton4),
		client.Subscribe(elevatorButton1))

	time.AfterFunc(1*time.Second, func() {
		mergedRequests.Close()
	})

	// send request to controller
	for it := range mergedRequests.Updates() {
		cs.PickUp(it.PickupFloor, it.Direction, it.FromDevice.Type, it.FromDevice.ID)
	}

	for i := 0; i < 5; i++ {
		cs.Step(1)
	}

}
