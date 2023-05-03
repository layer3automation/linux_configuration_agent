package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/layer3automation/linux_configuration_agent/configurator"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 8081))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterConfigurationAgentServiceServer(grpcServer, pb.NewServer())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("couldn't serve gRPC server")
	}
}
