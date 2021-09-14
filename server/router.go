package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true) //.StrictSlash(true) // .PathPrefix("api/v1").Subrouter()

	router.HandleFunc("/", HomePage)
	router.HandleFunc("/tasks", GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks", AddTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", UpdateTask).Methods("PUT")

	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"POST", "GET", "DELETE", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})

	// log.Fatal(http.ListenAndServe(":8090", router))
	log.Fatal(http.ListenAndServe(":8090", handlers.CORS(credentials, methods, origins)(router)))
}
