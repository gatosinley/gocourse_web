package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	port := ":3333"
	http.HandleFunc("/users", getCourses)
	http.HandleFunc("/courses", getCourses)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got user")
	io.WriteString(w, "this is my user endpoint")
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got Courses")
	io.WriteString(w, "this is my course endpoint")
}
