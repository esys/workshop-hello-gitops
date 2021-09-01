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
	_, err := os.Hostname()
	if err != nil {
		log.Printf("unknown host")
	}
<<<<<<< Updated upstream
	resp := fmt.Sprintf("Hello  to Salvatore and Zahoor from Piyush!",)

=======
	resp := fmt.Sprintf("Hello  to awesome VMware Tanzu team from Piyush!",)
>>>>>>> Stashed changes
	_, err = w.Write([]byte(resp))
	if err != nil {
		log.Panicf("not able to write http output: %s", err)
	}
}
