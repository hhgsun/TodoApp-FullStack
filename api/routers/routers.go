package routers

import (
	"log"
	"net/http"

	"goapi/controller"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true) //.StrictSlash(true) // .PathPrefix("api/v1").Subrouter()

	router.HandleFunc("/", controller.HomePage)
	router.HandleFunc("/tasks", controller.GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks", controller.AddTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", controller.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", controller.DeleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", controller.UpdateTask).Methods("PUT")

	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"POST", "GET", "DELETE", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(":8090", handlers.CORS(credentials, methods, origins)(router)))
}
