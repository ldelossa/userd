package main

import (
	"fmt"
	"net"
	"os"

	pb "github.com/ldelossa/userd/user"
	"google.golang.org/grpc"
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
	go h.ListenAndServe()

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := grpc.NewServer()
	pb.RegisterUserGRPCServer(s, &GRPCServer{ds: ds})
	fmt.Println("Starting grpc server")
	s.Serve(lis)
}
