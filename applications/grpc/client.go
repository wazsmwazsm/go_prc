package main

import (
	"fmt"
	"grpctest/user"

	pb "grpctest/user/proto"

	"context"

	"google.golang.org/grpc"
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
	fmt.Println("Get user list by stream")

	if res, err := GetStream(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	fmt.Println("Get user list  by stream successed")

	// test set
	fmt.Println("Set user")
	err = SetUserInfo("Qin", 27, "Beijing")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Set successed")

	// get get list
	fmt.Println("Get user list")

	if res, err := GetList(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func GetList() ([]user.User, error) {
	var list []user.User
	res, err := client.GetList(context.Background(), &pb.EmptyRequest{})
	if err != nil {
		return []user.User{}, err
	}

	for _, userResponse := range res.List {
		list = append(list, user.User{
			Id:      userResponse.GetId(),
			Name:    userResponse.GetName(),
			Age:     userResponse.GetAge(),
			Address: userResponse.GetAddress(),
		})
	}

	return list, nil
}

func GetStream() ([]user.User, error) {

	var list []user.User

	res, err := client.GetStream(context.Background(), &pb.EmptyRequest{})
	if err != nil {
		return []user.User{}, err
	}

	for {
		result, err := res.Recv()
		if err != nil {
			break
		}

		list = append(list, user.User{
			Id:      result.GetId(),
			Name:    result.GetName(),
			Age:     result.GetAge(),
			Address: result.GetAddress(),
		})

	}

	return list, nil
}

func GetUserInfo(id int32) (user.User, error) {
	result, err := client.Get(context.Background(), &pb.UserGetRequest{Id: id})
	if err != nil {
		fmt.Println("Request failed")
		return user.User{}, err
	}

	return user.User{
		Id:      result.GetId(),
		Name:    result.GetName(),
		Age:     result.GetAge(),
		Address: result.GetAddress(),
	}, nil
}

func SetUserInfo(name string, age int32, address string) error {
	_, err := client.Set(context.Background(),
		&pb.UserSetRequest{Name: name, Age: age, Address: address})
	if err != nil {
		fmt.Println("Request failed")
		return err
	}

	return nil
}
