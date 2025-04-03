package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal("Error making request:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response:", err)
	}

	fmt.Printf("Server response: %s\n", body)
}
