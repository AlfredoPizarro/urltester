package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	urls := []string{
		os.Getenv("SERVER_IP_URL"),
		os.Getenv("SERVER_SERVICE_URL"),
		os.Getenv("SERVER_ROUTE_URL"),
	}

	for _, url := range urls {
		if url == "" {
			continue
		}

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
