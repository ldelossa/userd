package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	uuid "github.com/google/uuid.git"
)

type Location struct {
	State       string  `json:"state,omitempty"`
	City        string  `json:"city,omitempty"`
	ZipCode     int     `json:"zipCode,omitempty"`
	PhoneNumber string  `json:"phoneNumber,omitempty"`
	Lat         float64 `json:"lat,omitempty"`
	Long        float64 `json:"long,omitempty"`
}

type User struct {
	ID         string   `json:"id,omitempty"`
	Username   string   `json:"username,omitempty"`
	Password   string   `json:"password,omitempty"`
	Email      string   `json:"email,omitempty"`
	FName      string   `json:"firstName,omitempty"`
	MName      string   `json:"middleName,omitempty"`
	LName      string   `json:"lastName,omitempty"`
	Location   Location `json:"location,omitempty"`
	Reputation int      `json:"reputation,omitempty"`
}

func NewUser() *User {
	return &User{ID: uuid.New().String()}
}

// Value implements Valueer interface to marshal object into []byte type before
// storing into DB.
func (u User) Value() (driver.Value, error) {
	j, err := json.Marshal(u)
	return j, err
}

// Scan implements Scanner interface to Unmarshal return []byte array from DB into User
func (u *User) Scan(src interface{}) error {
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
