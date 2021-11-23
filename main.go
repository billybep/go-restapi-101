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

	listActivities := []ToDo{}

	listActivities = append(listActivities, ToDo{"Learning RESTAPI with Golang", "2021-11-05"})
	listActivities = append(listActivities, ToDo{"Testing Endpoint GET", "2021-11-06"})

	// create endpoint
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		// GET /
		if r.Method == "GET" {
			rw.Header().Add("Content-Type", "application/json")

			res := JSONResponse{
				http.StatusOK,
				true,
				"List of activities success",
				listActivities,
			}

			// convert res to json
			dataJSON, err := json.Marshal(res)

			if err != nil {
				fmt.Println("Something Wrong")
				http.Error(rw, "Something wrong", http.StatusInternalServerError)
				return
			}

			rw.Write(dataJSON)
			return

		} else if r.Method == "POST" {

		}

	})

	fmt.Println("Listning on port: 8080 ....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}