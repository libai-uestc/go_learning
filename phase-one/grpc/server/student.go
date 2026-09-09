package main

import (
	"context"
	"fmt"
	"io"
	grpc_model "libai/go/basic/phase-one/grpc/idl/model"
	grpc_service "libai/go/basic/phase-one/grpc/idl/service"
)

type Student struct {
	grpc_service.UnimplementedStudentServer
}

// func (s Student) QueryStudent(ctx context.Context, query *grpc_service.QueryStudentRequest) (resp *grpc_service.QueryStudentResponse, err error) {
// 	fmt.Printf("request: %+v\n", query)
// 	resp = &grpc_service.QueryStudentResponse{
// 		Students: []*grpc_model.Student{
// 			{Id: 123, Name: "李白", Age: 18},
// 			{Id: 456, Name: "libai", Age: 28},
// 		},
// 	}
// 	return
// }

func (s Student) QueryStudent(ctx context.Context, query *grpc_service.QueryStudentRequest) (resp *grpc_service.QueryStudentResponse, err error) {
	fmt.Printf("收到客户端请求: %+v\n", query)

	// 1. 模拟数据库里的所有数据
	allStudents := []*grpc_model.Student{
		{Id: 123, Name: "李白", Age: 18},
		{Id: 456, Name: "libai", Age: 28},
	}

	// 2. 准备一个切片存放匹配的结果
	var matchedStudents []*grpc_model.Student

	// 3. 根据客户端传来的 query.Id 进行过滤
	for _, stu := range allStudents {
		if stu.Id == query.Id {
			matchedStudents = append(matchedStudents, stu)
		}
	}

	// 4. 返回过滤后的结果
	resp = &grpc_service.QueryStudentResponse{
		Students: matchedStudents,
	}
	return resp, nil
}

// Server streaming RPC
func (s Student) QueryStudents2(query *grpc_service.StudentIds, server grpc_service.Student_QueryStudents2Server) error {
	for i := 0; i < 2; i++ {
		id := int64(i) + 100
		stu := &grpc_model.Student{Id: id, Name: "李白", Age: 18}
		err := server.Send(stu) // 向流中发送一个结果
		if err != nil {
			fmt.Printf("send Student %d failed: %s\n", id, err)
			return err
		}
	}
	return nil
}

// Client streaming RPC
func (s Student) QueryStudents3(server grpc_service.Student_QueryStudents3Server) error {
	students := make([]*grpc_model.Student, 0, 10)
	for {
		sid, err := server.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("recv request3 failed: %s\n", err)
			continue
		}
		stu := &grpc_model.Student{Id: sid.Id, Name: "李白", Age: 18}
		students = append(students, stu)
	}
	return server.SendMsg(&grpc_service.QueryStudentResponse{Students: students})
}

// Bidirectional streaming RPC
func (s Student) QueryStudents4(server grpc_service.Student_QueryStudents4Server) error {
	for {
		sid, err := server.Recv() // 从流中取出一个结果
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("recv request3 failed: %s\n", err)
			continue
		}
		stu := &grpc_model.Student{Id: sid.Id, Name: "李白", Age: 18}
		err = server.Send(stu) // 向流中发送一个结果
		if err != nil {
			fmt.Printf("send Student %d failed: %s\n", stu.Id, err)
			return err
		}
	}
	return nil
}
