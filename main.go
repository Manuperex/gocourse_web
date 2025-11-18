package main

import (
	"log"
	"net/http"
	"time"

	"gocourse_web/internal/user"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	userEndpoints := user.MakeEndpoints()

	router.HandleFunc("/users", userEndpoints.Create).Methods("POST")
	// router.HandleFunc("/users", userEndpoints.Get).Methods("GET")
	router.HandleFunc("/users", userEndpoints.GetAll).Methods("GET")
	router.HandleFunc("/users", userEndpoints.Update).Methods("PATCH")
	router.HandleFunc("/users", userEndpoints.Delete).Methods("DELETE")

	srv := &http.Server{
		Handler: 	  router,
		Addr: 		  "127.0.0.1:8000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}