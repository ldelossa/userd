package main

import (
	"context"

	uuid "github.com/google/uuid.git"
	pb "github.com/ldelossa/userd/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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
		return nil, grpc.Errorf(codes.Aborted, err.Error())
	}
	return u, nil
}

func (gs *GRPCServer) GetUserByID(ctx context.Context, id *pb.ID) (*pb.User, error) {
	u, err := gs.ds.GetUserByID(*id)
	if err != nil {
		return nil, grpc.Errorf(codes.NotFound, err.Error())
	}
	return u, nil
}

func (gs *GRPCServer) DeleteUserByID(ctx context.Context, id *pb.ID) (*pb.Empty, error) {
	err := gs.ds.DeleteUserByID(*id)
	if err != nil {
		return nil, grpc.Errorf(codes.NotFound, err.Error())
	}
	return &pb.Empty{}, nil
}
