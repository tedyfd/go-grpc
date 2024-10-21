package main

import (
	"context"
	"log"

	pb "github.com/tedyfd/go-grpc/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 2,
	})

	if err != nil {
		log.Fatalf("Couldn't sum: %v\n", err)
	}

	log.Printf("Sum: %s\n", res.Result)
}
