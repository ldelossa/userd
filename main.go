package main

import (
	"fmt"
	"os"
)

func main() {
	connStr := "user=postgres dbname=userd password=dev host=localhost sslmode=disable"
	ds, err := NewPGDatastore(connStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	h := NewHTTPServer(ds, 8080)
	fmt.Println("Starting http server")
	err = h.ListenAndServe()
	fmt.Println(err)
}
