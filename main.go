package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gatosinley/gocourse_web/internal/user"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	router := mux.NewRouter()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"cbr_user",
		"4k39wGoVdM1U6A53",
		"10.100.0.8",
		"3306",
		"COCHA_BUSINESS_RULES")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error en DB: ", err.Error())
	}
	db = db.Debug()

	_ = db.AutoMigrate(&user.User{})

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
