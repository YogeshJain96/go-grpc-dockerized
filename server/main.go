package main

import (
	"context"
	pb "demo/server/pb/org"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedDemoServer
}

func (s *server) GetEmpList(ctx context.Context, in *pb.GetEmpListRequest) (*pb.GetEmpListResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	return &pb.GetEmpListResponse{
		Emps: getExampleEmployees(),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterDemoServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getExampleEmployees() []*pb.Emp {
	someEmployees := []*pb.Emp{
		{
			Id:   1,
			Name: "Elon",
		},
		{
			Id:   2,
			Name: "Jeff",
		},
	}
	return someEmployees
}
