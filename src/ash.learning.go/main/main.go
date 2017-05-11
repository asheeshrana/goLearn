package main

import (
	"io/ioutil"
	"os"

	"ash.learning.go/configuration"
	"ash.learning.go/hello"
	"ash.learning.go/log"
	"ash.learning.go/simulator/handlers"
	"ash.learning.go/simulator/server"
)

func main() {
	hello.Hello()

	initialize()
	start()
}

func initialize() {
	initConfig()
	initLogging()
	initHandlers()
}

func initConfig() {
	//args := os.Args
	/*if len(args) < 2 {
		panic("Configuration file not specified in the input")
	}*/

	//Initialize configuration
	//configuration.LoadConfiguration(args[1])
	configuration.LoadConfiguration("C:/Asheesh/temp/golearn/configuration.txt")

}

func initLogging() {
	log.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
}

func initHandlers() {
	server.ServeMultiplexer.HandleFunc("/", handlers.RootHandler)

}

func start() {
	server.StartServer()
}
