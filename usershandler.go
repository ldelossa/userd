package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	uuid "github.com/google/uuid.git"
	"github.com/gorilla/mux"

	pb "github.com/ldelossa/userd/user"
)

// HandleUsers handles operations on the http endpoint: /users for our service.
func (h *HTTPServer) HandleUsers(w http.ResponseWriter, r *http.Request) {
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
		u, err := h.ds.GetUserByID(pb.ID{Id: username})
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
		u := &pb.User{Location: &pb.User_Location{}}
		err := decoder.Decode(u)
		u.Id = uuid.New().String()
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
		err := h.ds.DeleteUserByID(pb.ID{Id: id})
		if err != nil {
			errMsg := fmt.Sprintf("Could not delete user: %s", err.Error())
			JsonError(w, &Response{Message: errMsg}, http.StatusBadRequest)
			return
		}
		return
	}
}
