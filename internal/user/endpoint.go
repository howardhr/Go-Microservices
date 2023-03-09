package user

import (
	"encoding/json"
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

func MakeEndpoints() Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(),
		Get:    makeGetEndpoint(),
		GetAll: makeGetAllEndpoint(),
		Update: makeUpdateEndpoint(),
		Delete: makeDeleteEndpoint(),
	}
}

func makeCreateEndpoint() Controller {
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
		json.NewEncoder(w).Encode(req)
	}
}
func makeGetEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Get User")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
func makeGetAllEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("GetAll User")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
func makeUpdateEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Update User")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
func makeDeleteEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Delete User")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
