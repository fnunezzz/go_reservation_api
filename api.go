package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	listenAddress string
	storage Storage
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

var router *mux.Router

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func NewServer(listenAddress string, storage Storage) *ApiServer {
	return &ApiServer{
		listenAddress: listenAddress,
		storage: storage,
	}
}

func (s *ApiServer) handlers() {
	s.UserHandler(router)
}

func (s *ApiServer) Run() {
	router = mux.NewRouter()
	
	s.handlers()

	log.Println("JSON API Server running on port: ", s.listenAddress)

	http.ListenAndServe(s.listenAddress, router)

}

