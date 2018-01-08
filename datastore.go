package main

import (
	pb "github.com/ldelossa/userd/user"
)

// DataStore interface provides the methods which our service expects a Datastore to
// support.
type DataStore interface {
	init() error
	AddUser(u *pb.User) error
	GetUserByID(ID pb.ID) (*pb.User, error)
	GetUserByUserName(username string) (*pb.User, error)
	DeleteUserByID(ID pb.ID) error
	DeleteUserByUserName(username string) error
}
