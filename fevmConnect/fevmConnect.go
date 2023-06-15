package main

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/core/node"
	"github.com/ipfs/kubo/plugin"
	"go.uber.org/fx"
)

var logger = logging.Logger("fevmConnect")

// Plugins sets the list of plugins to be loaded.
var Plugins = []plugin.Plugin{
	&nfevmConnectPlugin{},
}

type fevmConnectPlugin struct{}

var _ plugin.PluginFx = (*fevmConnectPlugin)(nil)

func (p *fevmConnectPlugin) Name() string {
	return "fevmConnectPlugin"
}

func (p *fevmConnectPlugin) Version() string {
	return "0.0.1"
}

func (p *fevmConnectPlugin) Init(env *plugin.Environment) error {
	return nil
}

/* code from https://github.com/ipfs-shipyard/nopfs/blob/master/nopfs-kubo-plugin/plugin.go


// MakeBlocker is a factory for the blocker so that it can be provided with Fx.
func MakeBlocker() (*nopfs.Blocker, error) {
	files, err := nopfs.GetDenylistFiles()
	if err != nil {
		return nil, err
	}

	return nopfs.NewBlocker(files)
}

// PathResolvers returns wrapped PathResolvers for Kubo.
func PathResolvers(fetchers node.FetchersIn, blocker *nopfs.Blocker) node.PathResolversOut {
	res := node.PathResolverConfig(fetchers)
	return node.PathResolversOut{
		IPLDPathResolver:   ipfs.WrapResolver(res.IPLDPathResolver, blocker),
		UnixFSPathResolver: ipfs.WrapResolver(res.UnixFSPathResolver, blocker),
	}
}

func (p *nopfsPlugin) Options(info core.FXNodeInfo) ([]fx.Option, error) {
	logging.SetLogLevel("nopfs", "INFO")
	logger.Info("Loading Nopfs plugin: content blocking")

	opts := append(
		info.FXOptions,
		fx.Provide(MakeBlocker),
		fx.Decorate(ipfs.WrapBlockService),
		fx.Decorate(ipfs.WrapNameSystem),
		fx.Decorate(PathResolvers),
	)
	return opts, nil
}

*/