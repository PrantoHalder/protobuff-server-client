package main

import (
	"context"
	"fmt"
	"log"

	pb "main.go/proto/hellowrold"

	"google.golang.org/grpc"
)


func main () {
	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	clint := pb.NewHellowroldClient(conn)
	res,err := clint.SayHello(context.Background(),&pb.HelloRequest{
		Name: "Pranto",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}