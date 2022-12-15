package main

import (
	"context"
	pb "demo/client/pb/org"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewDemoClient(conn)
	bookList, err := client.GetEmpList(context.Background(), &pb.GetEmpListRequest{})
	if err != nil {
		log.Fatalf("failed to get employees list: %v", err)
	}
	log.Printf("Employees list: %v", bookList)
}
