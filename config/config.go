package config

type FlatfileJSONConfigData struct {
	Location string `json:"location"`
}

type FlatfileJSONConfig struct {
	Config *FlatfileJSONConfigData
}

var flatfileConfig = FlatfileJSONConfig{}

func (config *FlatfileJSONConfig) CreateStructure() interface{} {
	return &FlatfileJSONConfigData{
		Location: "data.json",
	}
}

func (config *FlatfileJSONConfig) Set(data interface{}) {
	config.Config = data.(*FlatfileJSONConfigData)
}

func Get() *FlatfileJSONConfig {
	return &flatfileConfig
}
