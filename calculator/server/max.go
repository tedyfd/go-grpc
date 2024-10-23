package main

import (
	"io"
	"log"

	pb "github.com/tedyfd/go-grpc/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("max function was invoked with \n")

	var maximum int32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		if number := req.Number; number > maximum {
			maximum = number
			err := stream.Send(&pb.MaxResponse{
				Result: maximum,
			})
			log.Printf("Maximum Send: %d", maximum)

			if err != nil {
				log.Fatalf("Error while sending: %v", err)
			}

		}
	}
}
