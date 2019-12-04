package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func greeting(writer http.ResponseWriter, request *http.Request) {
	// vars := mux.Vars(request)
	// // wishes := vars["wishes"]        // URL parameter
	// say := request.FormValue("say") // Query paramater...

	// fmt.Println(vars)
	// fmt.Println("query params=", say)

	// var urlParam = fmt.Sprintf("Your URL param is %s, and Query param is %s", wishes, say)

	// _, err := writer.Write([]byte(urlParam))

	// Sampling the JSON to send as reponse ...
	type employee struct {
		Name   string  `json:"name"`
		Id     string  `json:"id"`
		Salary float64 `json:"salary"`
	}

	e := employee{
		Name:   "shan",
		Id:     "222",
		Salary: 3.444,
	}

	emp, err := json.Marshal(e)

	if err != nil {
		log.Fatal(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	fmt.Fprint(writer, string(emp))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{wishes}", greeting).Methods("GET")
	// r.HandleFunc("/hello/{wishes}", greeting).Queries("say", "{say}")
	http.ListenAndServe(":5000", handlers.CORS()(r))
}
