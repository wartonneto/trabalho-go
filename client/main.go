package main

import (
	"context"
	"log"

	"github.com/wartonneto/trabalho-go/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewHelloClient(conn)

	req := &pb.HelloRequest{
		Id:          "1",
		Title:       "Magico de OZ",
		Description: "Livro sobre o mágico",
		Value:       "13",
	}

	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(res)
}
