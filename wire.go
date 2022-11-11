//go:build wireinject
// +build wireinject

package quick

import (
	"github.com/google/wire"
	"go-micro.dev/v5/types"
	"go-micro.dev/v5/log"
	"go-micro.dev/v5/registry"
)

func newService(
	options *Options,
	cliConfig cli.Config,
	logConfig log.Config,
	registryConfig registry.Config,
) (*Service, error) {
	panic(wire.Build(
		ProvideService,
	))
}
