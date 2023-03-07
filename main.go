package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "main.go/proto/hellowrold"
	"main.go/services/hellowrold"
)


func main(){
	lis ,err := net.Listen("tcp",":8000")
	if err != nil {
		log.Fatalln(err)
	}
	gs := grpc.NewServer()
	pb.RegisterHellowroldServer(gs,hellowrold.Server{})
	if err := gs.Serve(lis);err != nil{
		log.Fatalln(err)
	}
}