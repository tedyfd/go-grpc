package main

import (
	"context"
	"log"
	"time"

	pb "github.com/tedyfd/go-grpc/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "ted"},
		{FirstName: "teds"},
		{FirstName: "tedf"},
		{FirstName: "tedy"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receive response LongGreet %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
