package main

import (
	"fmt"
	pb "github.com/maanasasubrahmanyam-sd/customeValTest/generated/generated"
	server "github.com/maanasasubrahmanyam-sd/customeValTest/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

// server is used to implement the server

const (
	port = ":50051"
)
func main() {
	// protoc   -I ./interfaces/  -I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.1.0  --go_out="plugins=grpc:./generated"    --validate_out="lang=go:./generated" ./interfaces/test_server/*.proto
	//coding/testgo# go build cassandratest && ./cassandratest -ni 10 -nq 1000000 -del=false -pno=10
	fmt.Println("Go protobuffer test")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSearchServiceServer(s, &server.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
