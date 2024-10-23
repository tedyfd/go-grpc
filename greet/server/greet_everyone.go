package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/tedyfd/go-grpc/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone function was invoked \n")
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
		}

		res := fmt.Sprintf("hello %s!\n", req.FirstName)
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client %v\n", err)
		}
	}
}
