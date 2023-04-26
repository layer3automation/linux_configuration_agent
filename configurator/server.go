package configurator

import (
	context "context"
	"log"
)

type ConfigurationAgentServer struct {
	UnimplementedConfigurationAgentServiceServer
}

func NewServer() *ConfigurationAgentServer {
	return &ConfigurationAgentServer{}
}

func (c *ConfigurationAgentServer) AssignIPToInterface(ctx context.Context, assignment *IPAssignment) (*Result, error) {
	log.Println("AssignIP Called")
	return nil, nil
}
