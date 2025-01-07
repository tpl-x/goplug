package main

import (
	"fmt"
	"github.com/hashicorp/go-plugin"
	"github.com/hashicorp/go-plugin/examples/grpc/shared"
	"github.com/tpl-x/goplug/internal/pkg/pluginopt"
	v1 "github.com/tpl-x/goplug/internal/pkg/plugins/myip/v1"
	"io"
	"log"
	"os"
	"os/exec"
)

func serve() error {
	// We're a host. Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: pluginopt.Handshake,
		Plugins:         pluginopt.PluginMap,
		Cmd:             exec.Command("sh", "-c", os.Getenv("KV_PLUGIN")),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolGRPC},
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		return err
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("my_ip_finder_grpc")
	if err != nil {
		return err
	}
	if ipFinder, ok := raw.(v1.MyIPFinder); ok {

	}
	return nil
}

func main() {
	// We don't want to see the plugin logs.
	log.SetOutput(io.Discard)

	if err := serve(); err != nil {
		fmt.Printf("error: %+v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
