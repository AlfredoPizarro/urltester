package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	currentTime := time.Now()
	response := fmt.Sprintf("Hello from %s at %s on %s", hostname, currentTime.Format("15:04:05"), currentTime.Format("2006-01-02"))
	fmt.Fprintf(w, response)
}
