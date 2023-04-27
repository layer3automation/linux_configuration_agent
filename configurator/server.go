package configurator

import (
	context "context"
	"log"

	utils "github.com/layer3automation/linux_configuration_agent/utils"
)

type ConfigurationAgentServer struct {
	UnimplementedConfigurationAgentServiceServer
}

func NewServer() *ConfigurationAgentServer {
	return &ConfigurationAgentServer{}
}

func (c *ConfigurationAgentServer) AssignIPToInterface(ctx context.Context, assignment *IPAssignment) (*Result, error) {
	log.Println("AssignIP Called")
	utils.ExecuteCommand("ip", "addr", "add", assignment.Ip, "dev", assignment.Interface)
	return &Result{
		Executed: true,
		Message:  "ip assigned correctly",
	}, nil
}

func (c *ConfigurationAgentServer) ConfigureNat(ctx context.Context, in *NatConfiguration) (*Result, error) {
	log.Printf("ConfigureNat Called: %+v", in)
	utils.ExecuteCommand("iptables", "-t", "nat", "-A", "PREROUTING", "-d", in.GetLocalNetwork(), "-j", "NETMAP", "--to", in.GetRemoteNetwork())
	utils.ExecuteCommand("iptables", "-t", "nat", "-A", "POSTROUTING", "-s", in.GetRemoteNetwork(), "-j", "NETMAP", "--to", in.GetLocalNetwork())
	utils.ExecuteCommand("iptables", "-t", "mangle", "-A PREROUTING", "!", "-i", in.GetLocalInterface(), "-j", "MARK", "--set-mark", in.GetPacketMark())
	utils.ExecuteCommand("ip", "route", "add", "default", "via", in.GetNextHop(), "table", in.GetTableNumber())
	utils.ExecuteCommand("ip", "rule", "add", "fwmark", in.GetPacketMark(), "lookup", in.GetTableNumber())
	return &Result{
		Executed: true,
		Message:  "nat configured correctly",
	}, nil
}

func (c *ConfigurationAgentServer) AddRoute(ctx context.Context, in *Route) (*Result, error) {
	log.Println("AddRoute Called")
	utils.ExecuteCommand("ip", "route", "add", in.GetDestinationNetwork(), "via", in.GetNextHop())
	return &Result{
		Executed: true,
		Message:  "route added correctly",
	}, nil
}
