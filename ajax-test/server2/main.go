package main

import (
	"html/template"
	"net/http"
	"os/exec"
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

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func enableCustomHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Headers", "Access-Token")
}

func test2(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	enableCustomHeaders(&w)
	accessKey := r.Header.Get("Access-Token")
	println(accessKey)
	if accessKey != "yobro123" {
		//w.WriteHeader(http.StatusForbidden)

	}
	cmd := exec.Command("date")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/test", test)
	mux.HandleFunc("/test2", test2)
	http.ListenAndServe(":8081", mux)
}
