package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/howardhr/Go-Microservices/internal/user"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	router := mux.NewRouter()
	_ = godotenv.Load()
	l := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"))

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = db.Debug()
	_ = db.AutoMigrate(&user.User{})
	userRepo := user.NewRepo(l, db)
	userSrv := user.NewService(l, userRepo)
	userEnd := user.MakeEndpoints(userSrv)
	router.HandleFunc("/users", userEnd.Create).Methods("POST")
	router.HandleFunc("/users", userEnd.GetAll).Methods("GET")
	router.HandleFunc("/users/{id}", userEnd.Get).Methods("GET")
	router.HandleFunc("/users/{id}", userEnd.Update).Methods("PATCH")
	router.HandleFunc("/users/{id}", userEnd.Delete).Methods("DELETE")
	srv := &http.Server{
		//http.TimeoutHandler(router, 3*time.Second, "Timeout"),
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	log.Println("Listening on..", "http://localhost:8080/")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
