package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	serverPort = flag.Int("port", 8180, "Port the server will listen on")
	serverHost = flag.String("host", "localhost", "Host server will listen on")
)

// Dispatcher is a structure that parses incoming requests and parses them
// to necessary parts of the application.
type Dispatcher struct {
	Name string
}

func (rh Dispatcher) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(r.URL.Path))
}

func main() {
	flag.Parse()

	serverAddr := strings.Join([]string{*serverHost, strconv.Itoa(*serverPort)}, ":")

	h := Dispatcher{
		Name: "Default",
	}

	s := &http.Server{
		Addr:    serverAddr,
		Handler: h,
	}
	fmt.Println("Starting server on", serverAddr)
	log.Fatal(s.ListenAndServe())
}
