package server

import (
	"net"
	"net/http"
	"strconv"

	"fmt"

	"ash.learning.go/configuration"
	"ash.learning.go/log"
	"ash.learning.go/messageHandling"
)

//StartServer starts the http server and waits for requests in the specified port
func StartServer() {
	var configuration = configuration.GetConfiguration()
	var ipAddress = configuration.ServerParams.IPAddress
	var port = configuration.ServerParams.Port
	serverMux := http.NewServeMux()
	var listener net.Listener
	var err error

	serverMux.HandleFunc("/", a)
	fmt.Printf("Port = %v\n", strconv.Itoa(port))
	//log.Info.Print("port = " + ":" + strconv.Itoa(port))

	listener, err = net.Listen("tcp", ipAddress+":"+strconv.Itoa(port))
	if err != nil {
		//var msgDetails = messageHandling.GetMessageDetails(messageHandling.UnableToStartListener)
		//log.Error.Println(msgDetails)
		panic(err)
	}

	err = http.Serve(listener, serverMux)
	if err != nil {
		var msgDetails = messageHandling.GetMessageDetails(messageHandling.UnableToStartListener)
		log.Error.Println(msgDetails)
		panic(err)
	}

	log.Info.Println("Server started")

	log.Info.Println("Closing server")
}

var a = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte{51, 51, 51})
}
