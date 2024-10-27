package main

import (
	"context"
	"log"

	pb "github.com/tedyfd/go-grpc/blog/proto"
)

func CreateBlog(c pb.BlogServiceClient) string {
	log.Println("---createBlog was invoked---")

	blog := &pb.Blog{
		AuthorId: "Tedy",
		Title:    "My First Blog",
		Content:  "My Content",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected Error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)
	return res.Id
}
