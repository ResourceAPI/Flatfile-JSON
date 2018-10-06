package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type FlatfileJSONStorage struct {
	Location string
	Data     map[string][]map[string]interface{}
}

// Initialize the storage.
func (storage *FlatfileJSONStorage) Initialize() error {
	var data map[string][]map[string]interface{}

	// TODO Read from config
	storage.Location = "data.json"

	if _, err := os.Stat(storage.Location); err == nil {
		bytes, _ := ioutil.ReadFile(storage.Location)
		json.Unmarshal(bytes, &data)
	}

	if data == nil {
		data = make(map[string][]map[string]interface{})
	}

	storage.Data = data

	return nil // TODO
}

// Start the storage. Must be a blocking call.
func (storage *FlatfileJSONStorage) Start() error {
	return nil // TODO
}

// Graceful stopping of the storage with a 30s timeout.
func (storage *FlatfileJSONStorage) Stop() error {
	return nil // TODO
}

// Retrieve resources.
func (storage *FlatfileJSONStorage) GetResources(resource string, filters []interface{}) ([]map[string]interface{}, error) {
	resources, ok := storage.Data[resource]

	if !ok {
		return make([]map[string]interface{}, 0), nil
	}

	// TODO Filters

	return resources, nil
}

// Create resources.
func (storage *FlatfileJSONStorage) CreateResources(resource string, data []map[string]interface{}) error {
	_, ok := storage.Data[resource]

	if !ok {
		storage.Data[resource] = data
	} else {
		storage.Data[resource] = append(storage.Data[resource], data...)
	}

	storage.Save()

	return nil
}

// Update resources.
func (storage *FlatfileJSONStorage) UpdateResources(resource string, data []map[string]interface{}, filters []interface{}) error {
	return nil // TODO
}

// Delete resources.
func (storage *FlatfileJSONStorage) DeleteResources(resource string, filters []interface{}) error {
	return nil // TODO
}

func (storage *FlatfileJSONStorage) Save() {
	if _, err := os.Stat(storage.Location); os.IsNotExist(err) {
		f, err := os.Create(storage.Location)
		if err != nil {
			fmt.Println("Failed to save: " + err.Error())
		}
		f.Close()
	}

	bytes, _ := json.Marshal(storage.Data)
	err := ioutil.WriteFile(storage.Location, bytes, 0666)
	if err != nil {
		fmt.Println("Failed to save: " + err.Error())
	}
}
