package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// AddServer maintains server state for all handlers.
type AddServer struct {
	addr     string
	log      *os.File
	requests uint
	GET      uint
	POST     uint
}

// AddRequest handles the addition of an add request.
type AddRequest struct {
	Nums []int `json:"nums"`
}

// NewAddServer creates a AddServer object that will be used to maintain server state.
func NewAddServer(addr string, log *os.File) *AddServer {
	return &AddServer{
		addr:     addr,
		log:      log,
		requests: 0,
		GET:      0,
		POST:     0,
	}
}

// RootHandler handles all incoming requests to /.
func (a *AddServer) RootHandler(w http.ResponseWriter, r *http.Request) {
	// Set the logfile.
	log.SetOutput(a.log)
	// Check to see if the request is an http GET and increment requests.
	if r.Method == "GET" {
		// Increment GET and total requests.
		a.GET++
		a.requests++
		// Write OK status.
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"Server\":\"Addition\"}"))
	} else if r.Method == "POST" {

	}
	log.Printf(" %s", r.Header)
	// Print server stats after each request.
	fmt.Printf("GET Requests to /: %d\n", a.GET)
	fmt.Printf("POST Requests to /: %d\n", a.POST)
	fmt.Printf("Total Requests to /: %d\n", a.requests)
}

func main() {
	// Create the logfile writer.
	logfile, err := os.OpenFile("math/var/log/add_micro.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	// Defer the closing of the log file until it is not needed.
	defer logfile.Close()
	// Set the logging output the file above.
	log.SetOutput(logfile)
	// Create the multiplexer that will handler requests.
	mux := http.NewServeMux()
	// Create the add server and set port, logfile information.
	add := NewAddServer(":8090", logfile)
	// Root handler for the addition server.
	mux.HandleFunc("/", add.RootHandler)
	// Listen and server web server.
	http.ListenAndServe(add.addr, mux)
}
