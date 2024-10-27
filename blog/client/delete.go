package main

import (
	"context"
	"log"

	pb "github.com/tedyfd/go-grpc/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("---deleteBlog was invoked---")

	blogId := &pb.BlogId{Id: id}
	_, err := c.Deleteblog(context.Background(), blogId)
	if err != nil {
		log.Printf("Error while reading: %v\n", err)
	}

	log.Printf("Blog was deleted! \n")
}
