package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/tedyfd/go-grpc/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while opening stream doMax: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{4, 2, 7, 8, 3, 5, 9, 10, 3, 12}

		for _, number := range numbers {
			log.Printf("Send number: %v\n", number)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})
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
				log.Printf("problem while reading server stream %v\n", err)
				break
			}

			log.Printf("Received new maximum: %d", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
