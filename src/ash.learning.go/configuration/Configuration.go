package configuration

import (
	"encoding/json"
	"io/ioutil"

	"ash.learning.go/messageHandling"
)

//Config represents the configuration type
type Config struct {
	ServerParams serverParams
	TemplatesDir string
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
		msgDetails := messageHandling.GetMessageDetails(messageHandling.UnableToLoadConfiguration)
		panic(msgDetails)
	}
	json.Unmarshal(raw, &configuration)
	return configuration
}

//GetConfiguration returns the configuration that was loaded
func GetConfiguration() Config {
	return configuration
}
