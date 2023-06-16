package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gatosinley/gocourse_web/internal/user"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	router := mux.NewRouter()
	_ = godotenv.Load()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"))

	fmt.Println("dsn")
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error en DB: ", err.Error())
	}
	db = db.Debug()

	_ = db.AutoMigrate(&user.User{})
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	userRepo := user.NewRepo(logger, db)
	userService := user.NewService(logger, userRepo)
	userEndpoint := user.MakeEndpoints(userService)

	router.HandleFunc("/users", userEndpoint.Create).Methods("POST")
	router.HandleFunc("/users/{id}", userEndpoint.Get).Methods("GET")
	router.HandleFunc("/users", userEndpoint.GetAll).Methods("GET")
	router.HandleFunc("/users/{id}", userEndpoint.Update).Methods("PATCH")
	router.HandleFunc("/users/{id}", userEndpoint.Delete).Methods("DELETE")

	server := &http.Server{
		Handler:      http.TimeoutHandler(router, time.Second*5, "Timeout"),
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

/* del .env
DATABASE_USER="cbr_user"
DATABASE_PASSWORD=4k39wGoVdM1U6A53
DATABASE_HOST=10.100.0.8
DATABASE_PORT=3306
DATABASE_NAME=COCHA_BUSINESS_RULES
*/
