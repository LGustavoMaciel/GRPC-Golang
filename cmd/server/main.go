package main

import (
	"context"
	"go-grpc/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct{
 pb.UnimplementedSendMessageServer	
}

func (service *Server) RequestMessage(ctx context.Context,req *pb.Request) (*pb.Response, error){
	log.Println("menssagem recebida: ", req.GetMessage())

	response := &pb.Response{
		Status: 1,
	}
	return response, nil
} 

func (service *Server) mustEmbedUnimplementedSendMessageServer() {}

func main(){
	grpcServer := grpc.NewServer()

	pb.RegisterSendMessageServer(grpcServer, &Server{})

	port := ":5000"
	log.Println("Server runing on port", port)
	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal(err)
	}

	grpc_error := grpcServer.Serve(listener)

	if  grpc_error != nil {
			log.Fatal(grpc_error)
		
	}
}