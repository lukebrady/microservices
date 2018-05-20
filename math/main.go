package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// MathServer maintains server state for all handlers.
type MathServer struct {
	addr     string
	log      *os.File
	requests uint
	GET      uint
	POST     uint
}

// NewMathServer creates a AddServer object that will be used to maintain server state.
func NewMathServer(addr string) *MathServer {
	return &MathServer{
		addr:     addr,
		log:      nil,
		requests: 0,
		GET:      0,
		POST:     0,
	}
}

// AddHandler takes the addition request and forwards it to the addition microservice.
func (m *MathServer) AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Send a post to the addition microservice with the request.
		resp, err := http.Post("http://localhost:8090", "text/json", r.Body)
		if err != nil {
			w.WriteHeader(resp.StatusCode)
		}
		defer resp.Body.Close()
		// If the status is OK, write header as http 200.
		w.WriteHeader(http.StatusOK)
		// Read the POST responses body.
		rBody, _ := ioutil.ReadAll(resp.Body)
		w.Write(rBody)
	} else {
		resp, err := http.Get("http://localhost:8090")
		if err != nil {
			w.WriteHeader(resp.StatusCode)
		}
		defer resp.Body.Close()
		w.WriteHeader(http.StatusOK)
		rBody, _ := ioutil.ReadAll(resp.Body)
		w.Write(rBody)
	}
}

// DivHandler takes a divison request and forwards it to the division microservice.
func (m *MathServer) DivHandler(w http.ResponseWriter, r *http.Request) {

}

// MultHandler takes a multiplication request and forwards it to the multiplication service.
func (m *MathServer) MultHandler(w http.ResponseWriter, r *http.Request) {

}

// SubHandler takes a subtraction request and forwards it to the subtraction service.
func (m *MathServer) SubHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	m := NewMathServer(":8089")
	mux := http.NewServeMux()
	mux.HandleFunc("/add", m.AddHandler)
	mux.HandleFunc("/div", m.DivHandler)
	mux.HandleFunc("/mult", m.MultHandler)
	mux.HandleFunc("/sub", m.SubHandler)
	fmt.Println("Starting math server on port 8089.")
	http.ListenAndServe(m.addr, mux)
}
