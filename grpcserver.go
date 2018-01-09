package main

import (
	"context"

	uuid "github.com/google/uuid.git"
	pb "github.com/ldelossa/userd/user"
)

type GRPCServer struct {
	ds DataStore
}

func NewGRPCServer(ds DataStore) *GRPCServer {
	return &GRPCServer{
		ds: ds,
	}
}

func (gs *GRPCServer) AddUser(ctx context.Context, u *pb.User) (*pb.User, error) {
	u.Id = uuid.New().String()
	err := gs.ds.AddUser(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (gs *GRPCServer) GetUserByID(ctx context.Context, id *pb.ID) (*pb.User, error) {
	u, err := gs.ds.GetUserByID(*id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (gs *GRPCServer) DeleteUserByID(ctx context.Context, id *pb.ID) (*pb.Response, error) {
	err := gs.ds.DeleteUserByID(*id)
	if err != nil {
		return &pb.Response{Id: id.Id, Success: false}, err
	}
	return &pb.Response{Id: id.Id, Success: true}, nil
}
