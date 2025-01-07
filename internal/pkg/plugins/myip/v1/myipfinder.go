package v1

import (
	"context"
	"github.com/hashicorp/go-plugin"
	v1 "github.com/tpl-x/goplug/api/plugins/myip/v1"
	"google.golang.org/grpc"
)

var _ plugin.GRPCPlugin = (*MyIpFindPlugin)(nil)

type MyIPFinder interface {
	GetMyIp(ip string) (location, region string, err error)
}

type MyIpFindPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl MyIPFinder
}

func (m *MyIpFindPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	v1.RegisterMyIpServiceServer(server, &MyIpFinderGrpcServer{Impl: m.Impl})
	return nil
}

func (m *MyIpFindPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &MyIpFinderGrpcClient{client: v1.NewMyIpServiceClient(conn)}, nil
}
