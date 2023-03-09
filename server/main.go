package main

import (
	pb "GRPC-TODO/proto"
	"GRPC-TODO/server/postgres"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type StoreServer struct {
	pb.UnimplementedStoreServiceServer
}

func (s *StoreServer) CreateStore(ctx context.Context, in *pb.Store) (*pb.Store, error) {
	log.Printf("Recived %v", in.GetName())

	store, err := postgres.CreateStore(&pb.Store{
		Id: 1,
		Name: in.Name,
		Description: in.Description,
		IsOpen: in.IsOpen,
		Addresses: in.Addresses,
	})

	if err != nil {
		return nil, err
	}

	return store, nil

}

func main() {

	lis, err := net.Listen("tcp", ":8001")

	if err != nil {
		log.Fatalf("failed connection %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterStoreServiceServer(s, &StoreServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed connection %v", err)
	}
}
