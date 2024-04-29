package main

import (
	"fmt"
	"log"

	pb "gitee.com/shirdonl/goWebActualCombat/chapter5/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8078", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial error: %v\n", err)
	}
	defer conn.Close()

	// 实例化 ProgrammerService
	client := pb.NewProgrammerServiceClient(conn)

	// 调用服务
	req := new(pb.Request)
	req.Name = "shirdon"
	resp, err := client.GetProgrammerInfo(context.Background(), req)
	if err != nil {
		log.Fatalf("resp error: %v\n", err)
	}

	fmt.Printf("Recevied: %v\n", resp)
}
