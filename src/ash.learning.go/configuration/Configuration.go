package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"ash.learning.go/errorHandling"
)

//Config represents the configuration type
type Config struct {
	ServerParams serverParams
}

//configuration represents the configuration object to be used within the project
var configuration Config

type serverParams struct {
	IPAddress string
	Port      int
}

//LoadConfiguration : Loads configuration from the file given in the input
func LoadConfiguration(fullConfigFileName string) Config {
	raw, err := ioutil.ReadFile(fullConfigFileName)
	if err != nil {
		msgDetails := errorHandling.GetMessageDetails(errorHandling.UnableToLoadConfiguration)
		panic(msgDetails)
	}
	json.Unmarshal(raw, &configuration)
	fmt.Printf("Configuration = %v", configuration)
	return configuration
}

//GetConfiguration returns the configuration that was loaded
func GetConfiguration() Config {
	return configuration
}
