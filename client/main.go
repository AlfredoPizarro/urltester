package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	urls := []string{
		"http://server-ip:8080",
		"http://server-service:8080",
		"http://server-route:8080",
	}

	for _, url := range urls {
		fmt.Printf("Trying URL: %s\n", url)
		response, err := http.Get(url)
		if err != nil {
			log.Println("Get:", err)
			continue
		}

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println("ReadAll:", err)
			continue
		}

		log.Printf("HTTP Code: %d\n", response.StatusCode)
		log.Printf("Response: %s\n", body)
		return
	}

	log.Println("Unable to connect to the server")
}
