package main

// DataStore interface provides the methods which our service expects a Datastore to
// support.
type DataStore interface {
	init() error
	AddUser(u *User) error
	GetUserByID(ID string) (*User, error)
	GetUserByUserName(username string) (*User, error)
	DeleteUserByID(ID string) error
	DeleteUserByUserName(username string) error
}
