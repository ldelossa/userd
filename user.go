package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	uuid "github.com/google/uuid.git"
	pb "github.com/ldelossa/userd/user"
)

// MetaUser wraps our protobuff definition. We are now able to add
// sql specific methods and extend our definition outside of the
// generated code.
type MetaUser struct {
	pb.User
}

// NewUser is the constructor for a new user struct. You should call this constructor
// to populate the unique ID field of a user. Do not create User literal structs
func NewMetaUser() *MetaUser {
	return &MetaUser{User: pb.User{Id: uuid.New().String()}}
}

// Value implements Valueer interface to marshal object into []byte type before
// storing into DB.
func (u MetaUser) Value() (driver.Value, error) {
	j, err := json.Marshal(u)
	return j, err
}

// Scan implements Scanner interface to Unmarshal return []byte array from DB into User
func (u *MetaUser) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	err := json.Unmarshal(source, u)
	if err != nil {
		return err
	}

	return nil
}
