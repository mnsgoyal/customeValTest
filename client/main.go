//the client part
package main

import (
	"context"
	pb "github.com/maanasasubrahmanyam-sd/customeValTest/generated/generated"
	"google.golang.org/grpc"
	"log"
	"time"
)

type Server struct{}

const (
	address     = "localhost:50051"
)
//protoc   -I ./interfaces/  -I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.1.0 -I ${GOPATH}/src --go_out="plugins=grpc:./generated" --validate_out="lang=go:./generated"  --proto_path=C:/Users/Rohit/go/src/github.com/maanasasubrahmanyam-sd/customeValTest/interfaces/test_server/test.proto ./interfaces/test_server/*.proto
func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSearchServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Search(ctx, &pb.SearchRequest{Query: "Protocol Buffer",EmailId: "alex.test@gmail.com",Name: "maanasa@gmail.com", UserId : "4"})
	if err != nil {
		log.Fatalf("could not execute search: %v", err)
	}
	//Lets validate the respose
	res := r.Validate()
	if res != nil {
		log.Fatalf("Response validation failed: %v", err)
	}

	//v, err := grpc_playground_validator.NewValidator()

	log.Printf("Greeting: %s", r.SearchResponse)
}