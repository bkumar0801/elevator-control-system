# elevator-control-system

1. cd $GOPATH/src/github.com/
2. git clone https://github.com/bkumar0801/elevator-control-system.git
3. cd elevator-control-system
4. go run main.go

Pickup request can be sent from UP/DOWN floor button and from button within elevator too.
If button within elevator is pressed, it means, same elevator will serve the request. In case of request coming in from different buttons on different floors, priority is calculated to serve the request.