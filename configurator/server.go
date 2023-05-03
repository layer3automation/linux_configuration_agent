package configurator

import (
	context "context"
	"log"

	utils "github.com/layer3automation/linux_configuration_agent/utils"
)

// ConfigurationAgentServer is the struct that implements the gRPC procedures.
type ConfigurationAgentServer struct {
	UnimplementedConfigurationAgentServiceServer
}

// NewServer returns a new ConfigurationAgentServer object.
func NewServer() *ConfigurationAgentServer {
	return &ConfigurationAgentServer{}
}

// AssignIPToInterface receives an IP and a Interface name and it configure the given IP on the given Interface.
func (c *ConfigurationAgentServer) AssignIPToInterface(_ context.Context, assignment *IPAssignment) (*Result, error) {
	log.Printf("AssignIP Called with %v as input", assignment)
	err := utils.ExecuteCommand("ip", "addr", "add", assignment.Ip, "dev", assignment.Interface)
	if err != nil {
		return &Result{
			Executed: false,
			Message:  "ip not assigned correctly",
		}, nil
	}
	return &Result{
		Executed: true,
		Message:  "ip assigned correctly",
	}, nil
}

// ConfigureNat configures nat with the given information.
func (c *ConfigurationAgentServer) ConfigureNat(_ context.Context, natConfiguration *NatConfiguration) (*Result, error) {
	log.Printf("ConfigureNat Called with %+v as input", natConfiguration)

	err := utils.ExecuteCommand("iptables", "-t", "nat", "-A", "PREROUTING",
		"-d", natConfiguration.GetLocalNetwork(), "-j", "NETMAP", "--to", natConfiguration.GetRemoteNetwork())
	if err != nil {
		return &Result{
			Executed: false,
			Message:  "nat prerouting not configured correctly",
		}, nil
	}

	err = utils.ExecuteCommand("iptables", "-t", "nat", "-A", "POSTROUTING",
		"-s", natConfiguration.GetRemoteNetwork(), "-j", "NETMAP", "--to", natConfiguration.GetLocalNetwork())
	if err != nil {
		return &Result{
			Executed: false,
			Message:  "nat postrouting not configured correctly",
		}, nil
	}

	err = utils.ExecuteCommand("iptables", "-t", "mangle", "-A PREROUTING",
		"!", "-i", natConfiguration.GetLocalInterface(), "-j", "MARK", "--set-mark", natConfiguration.GetPacketMark())
	if err != nil {
		return &Result{
			Executed: false,
			Message:  "mangle postrouting not configured correctly",
		}, nil
	}

	err = utils.ExecuteCommand("ip", "route", "add", "default",
		"via", natConfiguration.GetNextHop(), "table", natConfiguration.GetTableNumber())
	if err != nil {
		return &Result{
			Executed: false,
			Message:  "route not added correctly",
		}, nil
	}

	err = utils.ExecuteCommand("ip", "rule", "add", "fwmark",
		natConfiguration.GetPacketMark(), "lookup", natConfiguration.GetTableNumber())
	if err != nil {
		return &Result{
			Executed: false,
			Message:  "rule not added correctly",
		}, nil
	}

	return &Result{
		Executed: true,
		Message:  "nat configured correctly",
	}, nil
}

// AddRoute receives a destination network and a nexthop and configure a route.
func (c *ConfigurationAgentServer) AddRoute(_ context.Context, route *Route) (*Result, error) {
	log.Printf("AddRoute Called with %+v as input", route)
	err := utils.ExecuteCommand("ip", "route", "add", route.GetDestinationNetwork(), "via", route.GetNextHop())
	if err != nil {
		return &Result{
			Executed: false,
			Message:  "route not added correctly",
		}, nil
	}

	return &Result{
		Executed: true,
		Message:  "route added correctly",
	}, nil
}
