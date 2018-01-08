package main

import (
	"encoding/json"
	"fmt"

	pb "github.com/ldelossa/userd/user"
)

// This file contains template query strings. Out impementation
// injects marshaled json into the templates and return a final
// template string that can be used in query calls.

// SQL query templates, we inject marshaled Query structs
// into these templates
const (
	// Our implementation uses "ID", "Email", and "Username" as IDs.
	GetUserByIDTemp          = "SELECT u FROM users WHERE u @> '%s' OR u @> '%s' OR u @> '%s'"
	GetUserByUsernameTemp    = "SELECT u FROM users WHERE u @> '%s'"
	DeleteUserByIDTemp       = "DELETE FROM users WHERE u @> '%s' OR u @> '%s' OR u @> '%s'"
	DeleteUserByUsernameTemp = "DELETE FROM users WHERE u @> '%s'"
)

func NewGetUserByIDQuery(ID pb.ID) string {

	u, _ := json.Marshal(&pb.User{Username: ID.GetId()})
	e, _ := json.Marshal(&pb.User{Email: ID.GetId()})
	i, _ := json.Marshal(&pb.User{Id: ID.GetId()})

	return fmt.Sprintf(GetUserByIDTemp, u, e, i)
}

func NewGetUserByUsernameQuery(u string) string {
	un, _ := json.Marshal(&pb.User{Username: u})

	return fmt.Sprintf(GetUserByUsernameTemp, un)
}

func NewDeleteUserByIDQuery(ID pb.ID) string {
	u, _ := json.Marshal(&pb.User{Username: ID.GetId()})
	e, _ := json.Marshal(&pb.User{Email: ID.GetId()})
	i, _ := json.Marshal(&pb.User{Id: ID.GetId()})

	return fmt.Sprintf(DeleteUserByIDTemp, u, e, i)
}

func NewDeleteUserQuery(username string) string {
	m, _ := json.Marshal(&pb.User{Username: username})

	return fmt.Sprintf(DeleteUserByUsernameTemp, m)
}
