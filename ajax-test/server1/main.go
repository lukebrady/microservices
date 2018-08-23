package main

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	if err := tmp.Execute(w, nil); err != nil {
		panic(err)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Yo this is the /test route."))
}

func test2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Yo this is the /test2 route."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/test", test)
	mux.HandleFunc("/test2", test2)
	http.ListenAndServe(":8080", mux)
}
