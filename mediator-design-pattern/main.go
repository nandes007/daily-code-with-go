package main

func main() {
	stationManager := newStationManager()
	passangerTrain := &passangerTrain{
		mediator: stationManager,
	}
	goodsTrain := &goodsTrain{
		mediator: stationManager,
	}
	passangerTrain.requestArrival()
	goodsTrain.requestArrival()
	passangerTrain.departure()
}
