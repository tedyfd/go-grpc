package main

import (
	"io"
	"log"

	pb "github.com/tedyfd/go-grpc/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Printf("avg function was invoked with %v\n")

	var sum int32 = 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving number: %d", req.Number)
		sum += req.Number
		count++
	}
	return nil
}
