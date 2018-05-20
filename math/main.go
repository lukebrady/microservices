package main

import (
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

func (m *MathServer) AddHandler(w http.ResponseWriter, r *http.Request) {

}

func (m *MathServer) DivHandler(w http.ResponseWriter, r *http.Request) {

}

func (m *MathServer) MultHandler(w http.ResponseWriter, r *http.Request) {

}

func SubHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/add")
}
