package pluginopt

import "github.com/hashicorp/go-plugin"

var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "IP_FIND_PLUGIN",
	MagicCookieValue: "gopher",
}
