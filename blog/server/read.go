package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/tedyfd/go-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("ReadBlog was invoked with %v", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			fmt.Sprintf("Internal Error: %v", err),
		)
	}

	data := &BlogItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find blog with the ID",
		)
	}

	return documentToBlog(data), nil

}
