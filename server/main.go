package main

import (
	"context"
	example "github.com/mbier/example-grpc/.gen"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	example.RegisterPersonServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	example.UnimplementedPersonServiceServer
	personList []*example.Person
}

func (s *server) Save(ctx context.Context, req *example.CreateRequest) (*example.CreateResponse, error) {
	log.Printf("Save person: %v", req.GetPerson())

	s.personList = append(s.personList, req.GetPerson())

	return &example.CreateResponse{}, nil
}

func (s *server) List(ctx context.Context, req *example.ListRequest) (*example.ListResponse, error) {
	return &example.ListResponse{
		Person: s.personList,
	}, nil
}
