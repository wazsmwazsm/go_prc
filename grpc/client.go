package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpctest/user/proto"
	"grpctest/user"
)

var client pb.UserInfoServiceClient

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Connect err")
	}
	defer conn.Close()

	// create client
	client = pb.NewUserInfoServiceClient(conn)
	
	// test get 
	fmt.Println("Get user")
	
	if res, err := GetUserInfo(3); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	fmt.Println("Get successed")

	// test get list
	fmt.Println("Get user list")

	if res, err := GetList(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	fmt.Println("Get user list successed")

	// test set 
	fmt.Println("Set user")
	err = SetUserInfo("Qin", 27, "Beijing")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Set successed")

	// now get list
	fmt.Println("Now user list")

	if res, err := GetList(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func GetList() ([]user.User, error){

	var list []user.User

	res, err := client.GetList(context.Background(), &pb.EmptyRequest{})
	if err != nil {
		return []user.User{}, err
	}

	for {
        result, err := res.Recv()
        if err != nil {
            break
		}
		
        list = append(list, user.User{
			Id: result.GetId(), 
			Name: result.GetName(), 
			Age: result.GetAge(), 
			Address: result.GetAddress(),
		})
		
	}
	

	return list, nil
}

func GetUserInfo(id int32) (user.User, error) {
	result, err := client.Get(context.Background(), &pb.UserGetRequest{Id: id})
	if err != nil{
        fmt.Println("Request failed")
        return user.User{}, err
	}

	return user.User{
		Id: result.GetId(), 
		Name: result.GetName(), 
		Age: result.GetAge(), 
		Address: result.GetAddress(),
	}, nil
}

func SetUserInfo(name string, age int32, address string) error {
	_, err := client.Set(context.Background(), 
		&pb.UserSetRequest{Name: name, Age: age, Address: address})
	if err != nil{
        fmt.Println("Request failed")
        return err
	}

	return nil
}
