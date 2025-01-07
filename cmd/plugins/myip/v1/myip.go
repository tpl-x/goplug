package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/tpl-x/goplug/internal/pkg/pluginopt"
	v1 "github.com/tpl-x/goplug/internal/pkg/plugins/myip/v1"
)

var _ v1.MyIPFinder = (*myIpFinderImpl)(nil)

type myIpFinderImpl struct {
}

func (m myIpFinderImpl) GetMyIp() (ip, location, region string, err error) {
	return "", "", "", err
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginopt.Handshake,
		Plugins: map[string]plugin.Plugin{
			"my_ip_finder_grpc": &v1.MyIpFindPlugin{Impl: &myIpFinderImpl{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
