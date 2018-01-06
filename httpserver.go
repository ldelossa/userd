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

	r.HandleFunc("/users/{id}", h.HandleUsersEndpoint).Methods("GET", "DELETE")
	r.HandleFunc("/users", h.HandleUsersEndpoint).Methods("POST")
	return h
}

func (h *HTTPServer) HandleUsersEndpoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Parse username off pattern
		if _, ok := mux.Vars(r)["id"]; !ok {
			// http.Error(w, "missing username", http.StatusBadRequest)
			JsonError(w, &Response{Message: "missing id"}, http.StatusBadRequest)
			return
		}

		// Attempt user lookup
		username := mux.Vars(r)["id"]
		u, err := h.ds.GetUserByID(username)
		if err != nil {
			JsonError(w, &Response{Message: err.Error()}, http.StatusBadRequest)
			return
		}

		// return user json
		js, err := json.Marshal(u)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return

	case http.MethodPost:
		// Attempt to unmarshall request data into user struct
		decoder := json.NewDecoder(r.Body)
		u := NewUser()
		err := decoder.Decode(&u)
		if err != nil {
			errMsg := fmt.Sprintf("Json could not be unmarshaled. Error: %s", err.Error())
			JsonError(w, &Response{Message: errMsg}, http.StatusBadRequest)
			return
		}

		// Attempt to add user to datastore
		err = h.ds.AddUser(u)
		if err != nil {
			errMsg := fmt.Sprintf("Could not add user to datastore: %s", err.Error())
			JsonError(w, &Response{Message: errMsg}, http.StatusBadRequest)
			return
		}

		// Display 200 and created user as response
		j, _ := json.Marshal(u)
		w.Write(j)
		return

	case http.MethodDelete:
		// Parse username off patern
		if _, ok := mux.Vars(r)["id"]; !ok {
			JsonError(w, &Response{Message: "missing id"}, http.StatusBadRequest)
			return
		}

		// Attempt delete of user
		id := mux.Vars(r)["id"]
		err := h.ds.DeleteUserByID(id)
		if err != nil {
			errMsg := fmt.Sprintf("Could not delete user: %s", err.Error())
			JsonError(w, &Response{Message: errMsg}, http.StatusBadRequest)
			return
		}
		return
	}
}
