package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/users", getUsers)
	router.HandleFunc("/courses", getCourses)

	server := &http.Server{
		Handler:      http.TimeoutHandler(router, time.Second*5, "Timeout"),
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got user")
	time.Sleep(6 * time.Second)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got Courses")
	time.Sleep(4 * time.Second)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
