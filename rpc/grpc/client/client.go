package main

import (
	"context"
	"fmt"

	pb "github.com/iixlabs/virtual-lsobus/rpc/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9999", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	c := pb.NewChainAPIClient(conn)
	r, err := c.Version(context.Background(), &pb.VersionRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result, ", r)
}
