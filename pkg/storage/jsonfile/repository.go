package jsonfile

import (
	"encoding/json"
	"fga-asg-3/pkg/sensor"
	"log"
	"os"
)

type JsonFile struct {
	Status Status `json:"status"`
}

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Storage struct {
	jsonFile JsonFile
}

func NewStorage() *Storage {
	b, err := os.ReadFile("./web/static/status.json")

	if err != nil {
		log.Fatal(err)
	}

	jsonFile := JsonFile{}

	err = json.Unmarshal(b, &jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	return &Storage{
		jsonFile: jsonFile,
	}
}

func (s *Storage) GetStatus() (*sensor.SensorData, error) {
	return &sensor.SensorData{
		Status: sensor.Status{
			Water: s.jsonFile.Status.Water,
			Wind:  s.jsonFile.Status.Wind,
		},
	}, nil
}

func (s *Storage) SetStatus(data *sensor.SensorData) error {
	s.jsonFile.Status.Water = data.Status.Water
	s.jsonFile.Status.Wind = data.Status.Wind

	b, err := json.Marshal(s.jsonFile)
	if err != nil {
		return err
	}

	err = os.WriteFile("./web/static/status.json", b, 0644)
	if err != nil {
		return err
	}

	return nil
}
