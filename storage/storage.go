package storage

import (
	"encoding/json"
	"fmt"
	"github.com/StratoAPI/Interface/filter"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
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
func (storage *FlatfileJSONStorage) GetResources(resource string, filters []filter.ProcessedFilter) ([]map[string]interface{}, error) {
	resources, ok := storage.Data[resource]

	if !ok {
		return make([]map[string]interface{}, 0), nil
	}

	resultList := make([]map[string]interface{}, 0)
	for _, res := range resources {
		if resourceComplies(res, filters) {
			resultList = append(resultList, res)
		}
	}

	return resultList, nil
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
func (storage *FlatfileJSONStorage) UpdateResources(resource string, data []map[string]interface{}, filters []filter.ProcessedFilter) error {
	return nil // TODO
}

// Delete resources.
func (storage *FlatfileJSONStorage) DeleteResources(resource string, filters []filter.ProcessedFilter) error {
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

func resourceComplies(res map[string]interface{}, filters []filter.ProcessedFilter) bool {
	for _, f := range filters {
		switch f.Type {
		case "simple":
			casted, ok := f.Data.(*filter.Simple)

			if !ok {
				return false
			}

			data, found := resolveKey(res, strings.Split(casted.Key, "."))

			if !found {
				return false
			}

			reflect.ValueOf(data).Type().Kind()

			if casted.Operation != filter.OpEQ && casted.Operation != filter.OpNEQ {
				k := reflect.ValueOf(data).Type().Kind()
				if k == reflect.Invalid ||
					k == reflect.Bool ||
					k == reflect.Array ||
					k == reflect.Chan ||
					k == reflect.Func ||
					k == reflect.Interface ||
					k == reflect.Map ||
					k == reflect.Ptr ||
					k == reflect.Slice ||
					k == reflect.String ||
					k == reflect.Struct ||
					k == reflect.UnsafePointer {
					return false
				}
			}

			switch casted.Operation {
			case filter.OpEQ:
				if data != casted.Value {
					return false
				}
			case filter.OpNEQ:
				if data == casted.Value {
					return false
				}
			case filter.OpLT:
				fallthrough
			case filter.OpLTE:
				fallthrough
			case filter.OpGT:
				fallthrough
			case filter.OpGTE:
				dataFloat, err := getFloat(data)

				if err != nil {
					return false
				}

				valueFloat, err := getFloat(casted.Value)

				if err != nil {
					return false
				}

				switch casted.Operation {
				case filter.OpLT:
					if dataFloat >= valueFloat {
						return false
					}
				case filter.OpLTE:
					if dataFloat > valueFloat {
						return false
					}
				case filter.OpGT:
					if dataFloat <= valueFloat {
						return false
					}
				case filter.OpGTE:
					if dataFloat < valueFloat {
						return false
					}
				}
			}
		}
	}

	return true
}

func resolveKey(res map[string]interface{}, key []string) (interface{}, bool) {
	data, ok := res[key[0]]

	if !ok {
		return nil, false
	}

	if subMap, ok := data.(map[string]interface{}); ok {
		return resolveKey(subMap, key[1:])
	}

	return data, true
}

var floatType = reflect.TypeOf(float64(0))

func getFloat(unk interface{}) (float64, error) {
	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatType) {
		return 0, fmt.Errorf("cannot convert %v to float64", v.Type())
	}
	fv := v.Convert(floatType)
	return fv.Float(), nil
}
