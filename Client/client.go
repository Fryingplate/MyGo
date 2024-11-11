package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	url := "http://localhost:8080/"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Failed to make the request", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Server returned non 200 status:", resp.Status)
		os.Exit(1)
	}

	var responce Response

	if err := json.NewDecoder(resp.Body).Decode(&responce); err != nil {
		fmt.Println("Failed to decode the JSON ", err)
		os.Exit(1)
	}

	fmt.Println("Message ", responce.Message)
	fmt.Println("Status ", responce.Status)

}
