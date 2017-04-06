package multus

import (
	"encoding/json"
	"os"
)

// Config stores application config
type Config struct {
	DBString string `json:"db_string"`
}

// LoadConfig returns loaded config file
func LoadConfig(fileName string) Config {
	// open config file
	configFile, err := os.Open(fileName)
	if err != nil {
		panic("Could not open config file")
	}

	// decode json into our struct
	var config Config
	err = json.NewDecoder(configFile).Decode(&config)

	// check for error decoding config
	if err != nil {
		panic("Could not load config file")
	}

	return config
}