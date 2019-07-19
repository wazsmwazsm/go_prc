package main

import (
	"fmt"
	"google.golang.org/grpc"
	pb "grpctest/user/proto"
	"net"
	"grpctest/user"
)


func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("listen err")
		return
	}
	// create a grpc server
	gs := grpc.NewServer()
	pb.RegisterUserInfoServiceServer(gs, &user.UserService{})

	fmt.Println("Start serve...")
	gs.Serve(lis)
}
