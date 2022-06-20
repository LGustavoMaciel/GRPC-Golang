package main

import (
	"context"
	"go-grpc/pb"
	"log"

	"google.golang.org/grpc"
)

func main(){

	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewSendMessageClient(conn)

	req := &pb.Request{
		Message: "Hello GRPC",
	}

	res, err := client.RequestMessage(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	}
 log.Print(res.GetStatus())
}