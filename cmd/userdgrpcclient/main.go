package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ldelossa/userd/user"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	u := &pb.User{
		Email:    "grpcuser@grpc.com",
		Fname:    "GR",
		Lname:    "PC",
		Password: "grpc",
		Username: "grpc",
		Location: &pb.User_Location{
			City:        "Brooklyn",
			Phonenumber: "917-816-4857",
		}}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserGRPCClient(conn)

	uu, err := client.AddUser(context.Background(), u)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Created user:", uu)

	uu, err = client.GetUserByID(context.Background(), &pb.ID{Id: "grpc"})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Got user:", uu)

	_, err = client.DeleteUserByID(context.Background(), &pb.ID{Id: "grpc"})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Deleted user:", "grpc")

}
