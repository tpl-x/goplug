package main

import (
	"encoding/json"
	"github.com/hashicorp/go-plugin"
	"github.com/tpl-x/goplug/internal/pkg/pluginopt"
	v1 "github.com/tpl-x/goplug/internal/pkg/plugins/myip/v1"
	"log"
	"net/http"
	"os/exec"
)

type ipResponse struct {
	IPAddress string `json:"ip_address"`
	Location  string `json:"location"`
	Region    string `json:"region"`
}

func main() {
	// We're a host. Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: pluginopt.Handshake,
		Plugins:         pluginopt.PluginMap,
		Cmd:             exec.Command("plugins/findmyIp"),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolGRPC,
		},
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		panic(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("my_ip_finder_grpc")
	if err != nil {
		panic(err)
	}
	ipFinder, ok := raw.(v1.MyIPFinder)
	if !ok {
		panic("not a ip find plugin")
	}

	mux := http.NewServeMux()
	mux.Handle("GET /ip", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, location, region, err := ipFinder.GetMyIp()
		if err != nil {
			log.Println("failed to get ip", err)
			return
		}
		resp := ipResponse{
			IPAddress: ip,
			Location:  location,
			Region:    region,
		}
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Println("failed to encode to response", err)
			return
		}
	}))

}
