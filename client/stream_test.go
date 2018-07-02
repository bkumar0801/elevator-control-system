package client

import (
	"reflect"
	"testing"

	"github.com/mesosphere-elevator/elevator"
)

func TestSubscriptions(t *testing.T) {
	cases := []struct {
		device Component
		want   Request
	}{
		{
			device: &FloorButton{
				Pickup: Request{
					PickupFloor: 3,
					Direction:   elevator.UP,
					FromDevice: Device{
						Type: FB,
						ID:   3,
					},
				},
			},
			want: Request{
				PickupFloor: 3,
				Direction:   elevator.UP,
				FromDevice: Device{
					Type: FB,
					ID:   3,
				},
			},
		},
		{
			device: &FloorButton{
				Pickup: Request{
					PickupFloor: 4,
					Direction:   elevator.UP,
					FromDevice: Device{
						Type: FB,
						ID:   4,
					},
				},
			},
			want: Request{
				PickupFloor: 4,
				Direction:   elevator.UP,
				FromDevice: Device{
					Type: FB,
					ID:   4,
				},
			},
		},
	}
	for _, c := range cases {
		sub := Subscribe(c.device)
		got := <-sub.Updates()
		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("\n Updates mismatch:\n\texpected:%v\n\tgot:%v", c.want, got)
		}
		sub.Close()
	}
}
