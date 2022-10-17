package main

import (
	httpHandler "fga-asg-3/pkg/http"
	"fga-asg-3/pkg/sensor"
	"fga-asg-3/pkg/storage/jsonfile"
	"log"
	"net/http"
)

func main() {
	// Create a new sensor service
	sensorService := sensor.NewService(jsonfile.NewStorage())
	go sensorService.StartAutoRandomize()

	// Create a new http handler
	handler := httpHandler.NewHandler()
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
