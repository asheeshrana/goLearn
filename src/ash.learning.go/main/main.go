package main

import (
	"os"

	"ash.learning.go/configuration"
	"ash.learning.go/hello"
	"ash.learning.go/simulator/server"
)

func main() {
	hello.Hello()
	args := os.Args
	if len(args) < 2 {
		panic("Configuration file not specified in the input")
	}

	//Initialize configuration
	configuration.LoadConfiguration(args[1])

	server.StartServer()
}