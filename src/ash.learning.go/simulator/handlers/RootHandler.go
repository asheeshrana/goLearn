package handlers

import "net/http"

//RootHandler is provides default implementation for any incoming request
var RootHandler = func(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello there!!!"))
}
