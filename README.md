# Linux Configuration Agent

This project is based on the [template](https://github.com/layer3automation/template_configuration_agent) provided by the [layer3automation organization](https://github.com/orgs/layer3automation).

This implementation provides a simple way to configure a Linux machine used as a router. The functionalities that this implementation provides are:

* __AssignIPToInterface__
* __ConfigureNat__
* __AddRoute__

You can see the implementation of the three RPCs [here](configurator/server.go)

## How it works

The project starts a gRPC server and waits for external connections. The gRPC server must run on the same Linux machine that acts as a router.

The project uses `ip` command, which is installed in almost all Linux distributions by default.

### AssignIPToInterface

It accepts an object of type `IPAssignment`, which looks like

```go
type IPAssignment struct {
  Ip        string 
  Interface string 
}
```

and it runs the command `ip addr add <Assignment.IP> dev <Assignment.Interface>`

### ConfigureNat

It accepts an object of type `NatConfiguration`, which looks like

```go
type NatConfiguration struct {
  LocalNetwork   string
  RemoteNetwork  string
  LocalInterface string
  TableNumber    string
  PacketMark     string
  NextHop        string
}
```

and it runs the following commands

```bash
  iptables -t nat -A PREROUTING -d <NatConfiguration.LocalNetwork> -j NETMAP --to <NatConfiguration.RemoteNetwork>
  iptables -t nat -A POSTROUTING -s <NatConfiguration.RemoteNetwork> -j NETMAP --to <NatConfiguration.LocalNetwork>
  iptables -t mangle -A PREROUTING ! -i <NatConfiguration.LocalInterface> -j MARK --set-mark <NatConfiguration.PacketMark>
  ip route add default via <NatConfiguration.NextHop> table <NatConfiguration.TableNumber>
  ip rule add fwmark <NatConfiguration.PacketMark> lookup <NatConfiguration.TableNumber>
```

The nat is configured only on 1 side through `Policy-Based Routing`. A VRT is used and the number of the table to be used is received as a parameter.

### AddRoute

It accepts an object of type `Route` which looks like

```go
type Route struct {
  DestinationNetwork string
  NextHop            string
}
```

and it runs `ip route add <Route.DestinationNetwork> via <Route.NextHop>`

## How to use it

You can clone this repository and build the project by running the following commands:

```bash
git clone https://github.com/layer3automation/linux_configuration_agent.git
cd linux_configuration_agent
go build main.go
```

the last command will generate an executable in the project's folder that you can run. Be careful that the machine's architecture where the project is built must be the same as the linux router machine's one.
