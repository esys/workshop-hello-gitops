package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const PORT = 8080

func main() {
	startServer(handler)
}

func startServer(handler func(http.ResponseWriter, *http.Request)){
	http.HandleFunc("/", handler)
	log.Printf("starting server...")
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	log.Printf("received request from %s", r.Header.Get("User-Agent"))
	host, err := os.Hostname()
	if err != nil {
		host = "unknown host"
	}
	resp := fmt.Sprintf("Hello from %s", host)
	_, err = w.Write([]byte(resp))
	if err != nil {
		log.Panicf("not able to write http output: %s", err)
	}
}