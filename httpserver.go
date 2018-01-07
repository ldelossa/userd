package main

import (
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

// JsonError works like http.Error but uses our response
// struct as the body of the response. Like http.Error
// you will still need to call a naked return in the http handler
func JsonError(w http.ResponseWriter, r *Response, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	b, _ := json.Marshal(r)

	w.Write(b)
}

type HTTPServer struct {
	http.Server
	ds DataStore
}

func NewHTTPServer(ds DataStore, port int) *HTTPServer {
	r := mux.NewRouter()

	h := &HTTPServer{
		ds: ds,
		Server: http.Server{
			Addr:    fmt.Sprintf("localhost:%d", port),
			Handler: r,
		},
	}

	r.HandleFunc("/users/{id}", h.UsersHandler).Methods("GET", "DELETE")
	r.HandleFunc("/users", h.UsersHandler).Methods("POST")
	return h
}
