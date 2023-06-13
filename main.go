package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gatosinley/gocourse_web/internal/user"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	userService := user.NewService()
	userEndpoint := user.MakeEndpoints(userService)

	router.HandleFunc("/users", userEndpoint.Create).Methods("POST")
	//router.HandleFunc("/users", userEndpoint.Get).Methods("GET")
	router.HandleFunc("/users", userEndpoint.GetAll).Methods("GET")
	router.HandleFunc("/users", userEndpoint.Update).Methods("PATCH")
	router.HandleFunc("/users", userEndpoint.Delete).Methods("DELETE")

	server := &http.Server{
		Handler:      http.TimeoutHandler(router, time.Second*5, "Timeout"),
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
