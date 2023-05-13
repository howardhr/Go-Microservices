package main

import (
	"github.com/gorilla/mux"
	"github.com/howardhr/Go-Microservices/internal/course"
	"github.com/howardhr/Go-Microservices/internal/enrrolment"
	"github.com/howardhr/Go-Microservices/internal/user"
	"github.com/howardhr/Go-Microservices/pkg/bootstrap"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()
	_ = godotenv.Load()
	l := bootstrap.InitLoger()
	db, err := bootstrap.DBConnection()
	if err != nil {
		l.Fatalln(err)
	}
	userRepo := user.NewRepo(l, db)
	userSrv := user.NewService(l, userRepo)
	userEnd := user.MakeEndpoints(userSrv)

	courseRepo := course.NewRepo(db, l)
	courseSrv := course.NewService(l, courseRepo)
	courseEnd := course.MakeEndpoints(courseSrv)

	enrollRepo := enrrolment.NewRepo(db, l)
	enrollSrv := enrrolment.NewService(l, userSrv, courseSrv, enrollRepo)
	enrollEnd := enrrolment.MakeEndpoints(enrollSrv)

	router.HandleFunc("/users", userEnd.Create).Methods("POST")
	router.HandleFunc("/users", userEnd.GetAll).Methods("GET")
	router.HandleFunc("/users/{id}", userEnd.Get).Methods("GET")
	router.HandleFunc("/users/{id}", userEnd.Update).Methods("PATCH")
	router.HandleFunc("/users/{id}", userEnd.Delete).Methods("DELETE")

	router.HandleFunc("/courses", courseEnd.Create).Methods("POST")
	router.HandleFunc("/courses", courseEnd.GetAll).Methods("GET")
	router.HandleFunc("/courses/{id}", courseEnd.Get).Methods("GET")
	router.HandleFunc("/courses/{id}", courseEnd.Update).Methods("PATCH")
	router.HandleFunc("/courses/{id}", courseEnd.Delete).Methods("DELETE")

	router.HandleFunc("/enrrol", enrollEnd.Create).Methods("POST")

	srv := &http.Server{
		//http.TimeoutHandler(router, 3*time.Second, "Timeout"),
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	log.Println("Listening on..", "http://localhost:8080/")

	l.Fatalln(srv.ListenAndServe())

}
