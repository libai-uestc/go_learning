package main

import (
	"context"
	"fmt"
	user_service "libai/go/basic/phase-one/grpc/live/demo/idl/service"
	"log"
	"time"

	// "time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	// "google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("data/server.crt", "")
	if err != nil {
		panic(err)
	}
	conn, err := grpc.NewClient("localhost:2345",
		// grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTransportCredentials(creds),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(1<<20),
			grpc.MaxCallSendMsgSize(4<<20),
		),
		grpc.WithChainUnaryInterceptor(attchApiKey),
	)
	if err != nil {
		panic(err)
	}
	client := user_service.NewUserClient(conn)
	for i := 0; i < 3; i++ {
		begin := time.Now()
		ctx := context.Background()
		// ctx2, cancel := context.WithTimeout(ctx, 20000*time.Millisecond)
		// defer cancel()
		resp, err := client.Login(ctx, &user_service.LoginRequest{
			Name:     "李白",
			Password: "123456",
		},
		)
		if err != nil {
			log.Printf("grpc调用失败: %s", err)
		} else {
			fmt.Println(resp.Code, resp.Msg, "耗时", time.Since(begin).Milliseconds(), "ms")
		}
	}
}

// client 拦截器
func attchApiKey(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx = metadata.AppendToOutgoingContext(ctx, "api_key", "sk-123456")
	return invoker(ctx, method, req, reply, cc, opts...)
}

// http 公共的参数 放到 请求头里面 https (TLS)
// grpc(http2) 公共参数 请求头（metadata）
