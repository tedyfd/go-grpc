package main

import (
	"context"
	"log"

	pb "github.com/tedyfd/go-grpc/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Avg: %v\n", err)
	}

	number := []int32{3, 5, 9, 54, 23}

	for _, number := range number {
		log.Printf("sending number: %d\n", number)

		stream.Send(&pb.AvgRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}

	log.Printf("Avg: %f\n", res.Result)
}
