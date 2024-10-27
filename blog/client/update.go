package main

import (
	"context"
	"log"

	pb "github.com/tedyfd/go-grpc/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "update author",
		Title:    "update title",
		Content:  "update content",
	}
	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Printf("Error while updating: %v\n", err)
	}

	log.Printf("Blog was updated \n")
}
