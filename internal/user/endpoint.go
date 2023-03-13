package user

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)

	Endpoints struct {
		Create Controller
		Get    Controller
		GetAll Controller
		Update Controller
		Delete Controller
	}

	CreateReq struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}
	ErrorRes struct {
		Error string `json:"error"`
	}
)

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		Get:    makeGetEndpoint(s),
		GetAll: makeGetAllEndpoint(s),
		Update: makeUpdateEndpoint(s),
		Delete: makeDeleteEndpoint(s),
	}
}

func makeCreateEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRes{"Invalid Request format"})
			return
		}
		if len(req.FirstName) < 1 {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRes{"First Name is required"})
			return
		}
		if len(req.LastName) < 1 {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRes{"Last Name is required"})
			return
		}
		user, err := s.Create(req.FirstName, req.LastName, req.Email, req.Phone)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRes{err.Error()})
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}
func makeGetAllEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := s.GetAll()
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRes{err.Error()})
			return
		}

		json.NewEncoder(w).Encode(users)
	}
}
func makeGetEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		path := mux.Vars(r)
		id := path["id"]
		user, err := s.Get(id)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRes{"El usuario no existe"})
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}
func makeUpdateEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Update User")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
func makeDeleteEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Delete User")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
