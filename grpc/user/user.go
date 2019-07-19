package user

import (
	"golang.org/x/net/context"
	pb "grpctest/user/proto"
)

type User struct {
	Id      int32
	Name    string
	Age     int32
	Address string
}

var users = map[int32]User{
	1: {1, "jack", 18, "Beijing"},
	2: {2, "bob", 18, "Shanghai"},
	3: {3, "mark", 18, "Hangzhou"},
	4: {4, "jerry", 18, "New York"},
}

type UserService struct {

}

func (u *UserService) Get(ctx context.Context, in *pb.UserGetRequest) (*pb.UserGetResponse, error) {
	user := users[in.GetId()]
	return &pb.UserGetResponse{
		Id:      user.Id,
		Name:    user.Name,
		Age:     user.Age,
		Address: user.Address,
	}, nil
}

func (u * UserService) GetList(in *pb.EmptyRequest, res pb.UserInfoService_GetListServer) (error) {

	for _, user := range users {
		res.Send(&pb.UserGetResponse{
			Id:      user.Id,
			Name:    user.Name,
			Age:     user.Age,
			Address: user.Address,
		})
	}

	return nil
}

func (u *UserService) Set(ctx context.Context, in *pb.UserSetRequest) (*pb.UserSetResponse, error) {
	id := int32(len(users) + 1)
	users[id] = User{id, in.GetName(), in.GetAge(), in.GetAddress()}

	return &pb.UserSetResponse{Result: true}, nil
}
