package main

import (
	"database/sql"
	"errors"

	"fmt"

	"github.com/lib/pq"
)

// PGDataStore implements the DataStore interface.
// Embeds the sql.DB type to act as a DB instance
// We make use of jsonb fieds
type PGDataStore struct {
	connString string
	*sql.DB
}

func NewPGDatastore(connString string) (*PGDataStore, error) {
	pgd := &PGDataStore{
		connString: connString,
	}

	err := pgd.init()
	if err != nil {
		return nil, err
	}

	return pgd, nil
}

func (pgd *PGDataStore) init() error {
	// call Open to return init'd DB object
	var err error
	pgd.DB, err = sql.Open("postgres", pgd.connString)
	if err != nil {
		return err
	}

	// Determine if db can be reached.
	err = pgd.Ping()
	if err != nil {
		return err
	}

	return nil
}

// AddUser takes a user object and creates a json document in the
// configured postgresql database. The user object implements the
// Valuer interface. This interface seralizes the User object into
// a byte array before placing into the database.
func (pgd *PGDataStore) AddUser(u *User) error {
	// Issue insert and handle possible errors
	_, err := pgd.Exec("INSERT INTO users (u) VALUES ($1);", u)
	if err != nil {
		if pError, ok := err.(*pq.Error); ok {
			switch pError.Code {
			case "23505":
				errMsg := fmt.Sprintf("field already exists: %s", pError.Constraint)
				return errors.New(errMsg)
			case "23514":
				errMsg := fmt.Sprintf("field is missing: %s", pError.Constraint)
				return errors.New(errMsg)
			}
		}
		return err
	}

	return nil
}

// DeleteUserByUsername deletes a user from the configured postgresql database.
func (pgd *PGDataStore) DeleteUserByID(ID string) error {
	// Issue DELETE and handle possible errors
	res, err := pgd.Exec(NewDeleteUserByIDQuery(ID))
	affectedRows, _ := res.RowsAffected()
	switch {
	case affectedRows == 0:
		return errors.New("User does not exist in database. No delete action taken")
	case err != nil:
		return errors.New(fmt.Sprintf("Unhandled error: %s", err.Error()))
	default:
		return nil
	}
}

// DeleteUserByUsername deletes a user from the configured postgresql database.
func (pgd *PGDataStore) DeleteUserByUserName(username string) error {
	// Issue DELETE and handle possible errors
	res, err := pgd.Exec(NewDeleteUserQuery(username))
	affectedRows, _ := res.RowsAffected()
	switch {
	case affectedRows == 0:
		return errors.New("User does not exist in database. No delete action taken")
	case err != nil:
		return errors.New(fmt.Sprintf("Unhandled error: %s", err.Error()))
	default:
		return nil
	}
}

// GetUserByID retrieves a user by ID, Username, or Email. These fields are expected to be
// Unique and to only return a single row/document when queries for.
func (pgd *PGDataStore) GetUserByID(ID string) (*User, error) {
	u := &User{}
	err := pgd.QueryRow(NewGetUserByIDQuery(ID)).Scan(u)
	switch {
	case err == sql.ErrNoRows:
		return nil, errors.New("User does not exist in database")
	case err != nil:
		return nil, err
	default:
		return u, nil
	}
}

// GetUserByUsername retrieves a user by Username from the configured postgresql database.
func (pgd *PGDataStore) GetUserByUserName(username string) (*User, error) {
	u := &User{}
	query := fmt.Sprintf("SELECT u FROM users WHERE u @> '{\"username\":\"%s\"}'", username)
	err := pgd.QueryRow(query).Scan(u)
	switch {
	case err == sql.ErrNoRows:
		return nil, errors.New("User does not exist in database")
	case err != nil:
		return nil, err
	default:
		return u, nil
	}
}
