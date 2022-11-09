package main

import (
	"context"
	"log"
	"net"

	"github.com/wartonneto/trabalho-go/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHelloServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Nome do livro inserido: " + in.GetTitle()}, nil
}

func main() {
	println("Rodando servidor gRPC")

	listener, errGRPC := net.Listen("tcp", "127.0.0.1:9000")

	if errGRPC != nil {
		panic(errGRPC)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &Server{})
	if errGRPC := s.Serve(listener); errGRPC != nil {
		log.Fatalf("Erro fatal ao criar servidor gRPC %v", errGRPC)
	}
}
