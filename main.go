package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/layer3automation/linux_configuration_agent/configurator"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 8081))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterConfigurationAgentServiceServer(grpcServer, pb.NewServer())
	grpcServer.Serve(lis)
}
