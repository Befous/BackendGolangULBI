package main

import (
	"fmt"
	"log"
	"net/http"

	befous "github.com/Befous/BackendGolangULBI"
)

func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, befous.PostNear("publickey", "mongoenv", "befous", "geojson", r))
}

func handlerRequests() {
	http.HandleFunc("/", HelloHTTP)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handlerRequests()
}
