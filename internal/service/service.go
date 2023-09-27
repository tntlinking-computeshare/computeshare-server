package service

import (
	"github.com/google/wire"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/mohaijiang/computeshare-server/internal/conf"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewGreeterService,
	NewAgentService,
	NewStorageService,
	NewUserService,
	NewComputeInstanceService,
	NewComputePowerService,
	NewCronJob,
	NewIpfShell,
)

func NewIpfShell(c *conf.Data) *shell.Shell {
	return shell.NewShell(c.Ipfs.Url)
}
