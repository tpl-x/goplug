package pluginopt

import (
	"github.com/hashicorp/go-plugin"
	v1 "github.com/tpl-x/goplug/internal/pkg/plugins/myip/v1"
)

var PluginMap = map[string]plugin.Plugin{
	"my_ip_finder_grpc": &v1.MyIpFindPlugin{},
}
