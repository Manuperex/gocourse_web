package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)
	Endpoints  struct {
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
		Phone    string `json:"phone"`
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
			return
		}
		fmt.Println("Create user:", req)
		json.NewEncoder(w).Encode(req)
	}
}

func makeGetEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeGetAllEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeUpdateEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeDeleteEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}