package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Create Object TODO
type ToDo struct {
	Activity string `json:"activity"`
	Time string `json:"waktu"`
}

type JSONResponse struct {
	Code int `json:"code"`
	Success bool `json:"success"`
	Message string `json:"message"`
	Data []ToDo `json:"data"`
}

func main() {

	// create endpoint
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		// GET /
		if r.Method == "GET" {
			rw.Header().Add("Content-Type", "application/json")

			res := JSONResponse{
				http.StatusOK,
				true,
				"Testing Get Method",
				[]ToDo{},
			}

			// convert res to json
			dataJSON, err := json.Marshal(res)

			if err != nil {
				http.Error(rw, "Something wrong", http.StatusInternalServerError)
			}

			rw.Write(dataJSON)

		} else if r.Method == "POST" {

		}

	})

	fmt.Println("Listning on port: 8080 ....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}