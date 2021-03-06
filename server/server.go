package server

import (
	"context"
	pb "github.com/mnsgoyal/customeValTest/generated"
	log "github.com/sirupsen/logrus"
)

// server is used to implement the server
type Server struct{}

func (s *Server)Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchResponse, error){
	log.Printf("Received Emailid: %v", in.EmailId)
	log.Printf("Received Query: %v", in.Query)
	log.Printf("Received id: %v", in.Id)
	log.Printf("Received number: %v", in.Number)

	// Note this is the only place we use validate
	err := in.Validate()
	if err != nil {
		log.Warn("SearchRequest validation failed: %v", err)
		response := pb.SearchResponse{SearchResponse: "Invalid Inputs" }
		response.Validate()
		return &response, nil
	}else {
		response := pb.SearchResponse{SearchResponse: "Some Valid response from server " }
		response.Validate()
		return &response, nil
	}
}



