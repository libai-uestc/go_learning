package main

import (
	"context"
	"errors"
	"fmt"
	user_service "libai/go/basic/phase-one/grpc/live/demo/idl/service"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type MyUser struct {
	name                                 string
	user_service.UnimplementedUserServer // 匿名成员,"继承"
}

func (u *MyUser) Login(ctx context.Context, request *user_service.LoginRequest) (*user_service.LoginResponse, error) {
	u.name = "libai"
	fmt.Println(request.Name, request.Password)
	return &user_service.LoginResponse{
		Code: 234,
		Msg:  "登录成功",
	}, nil
}

func checkApiKey(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("获取不到IncomingContext")
	}
	value, exists := md["api_key"]
	if !exists {
		return nil, errors.New("获取不到api_key")
	}
	apiKey := value[0]
	// 检查apikey是否在白名单里面
	if apiKey != "sk-123456" {
		return nil, errors.New("非法的api_key")
	}
	return handler(ctx, req)
}

func main() {
	creds, err := credentials.NewServerTLSFromFile("data/server.crt", "data/rsa_private_key.pem")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(checkApiKey),
		grpc.Creds(creds),
	)
	user_service.RegisterUserServer(server, &MyUser{})
	listener, err := net.Listen("tcp", "127.0.0.1:2345")
	if err != nil {
		panic(err)
	}
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
