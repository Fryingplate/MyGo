package main

import (
	"encoding/json"
	"fmt"

	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func helloGO(w http.ResponseWriter, r *http.Request) {
	// create a instance of the response

	responce := Response{
		Message: "Hello  ichigo i am hereeeeeeeee",
		Status:  "200",
	}

	w.Header().Set("content-type", "application/json")

	if err := json.NewEncoder(w).Encode(responce); err != nil {
		http.Error(w, "Failer to encode the json ", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", helloGO)
	fmt.Println("Your server is starting on the port : 8080   ")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Sorry the server couldnt start: ", err)
	}
}
