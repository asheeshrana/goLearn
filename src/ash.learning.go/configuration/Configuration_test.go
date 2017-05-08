package configuration

import (
	"path/filepath"
	"testing"
)

//TestConfigurationLoad tests whether the configuration is loaded correctly or not
func TestConfigurationLoad(t *testing.T) {
	var configFile = filepath.Join("c:", "/Asheesh", "temp", "goLearn", "configuration.txt")
	t.Logf("Configuration file = %s", configFile)
	//Check with proper file
	var config = LoadConfiguration(configFile)

	var message string
	var failFlag = false
	//Check for each of your configuration below
	if config.ServerParams.IPAddress != "localhost" {
		message = message + "config.serverParams.ipAddress not defined correctly, "
		failFlag = true
	}

	if config.ServerParams.Port != 7777 {
		message = message + "config.serverParams.port not defined correctly, "
		failFlag = true
	}

	if failFlag {
		t.Log(message)
		t.Fail()
	}
}

func TestConfigurationLoadFail(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Received panic as expected: %v", r)
		} else {
			t.Fail()
		}
	}()
	var configFile = filepath.Join("c:", "/Asheesh", "temp", "goLearn", "configuration.txtd")
	//Check with proper file
	LoadConfiguration(configFile)
	t.Fail()
}
