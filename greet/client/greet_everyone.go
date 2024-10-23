package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/tedyfd/go-grpc/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "ted"},
		{FirstName: "teds"},
		{FirstName: "tedf"},
		{FirstName: "tedy"},
	}

	waitc := make(chan chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send Reqeust: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
