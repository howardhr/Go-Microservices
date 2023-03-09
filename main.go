package main

import (
	"github.com/gorilla/mux"
	"github.com/howardhr/Go-Microservices/internal/user"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()
	userEnd := user.MakeEndpoints()
	router.HandleFunc("/users", userEnd.Create).Methods("POST")
	router.HandleFunc("/users", userEnd.GetAll).Methods("GET")
	router.HandleFunc("/users", userEnd.Update).Methods("PATCH")
	router.HandleFunc("/users", userEnd.Delete).Methods("DELETE")
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
