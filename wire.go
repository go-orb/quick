//go:build wireinject
// +build wireinject

package quick

import (
	"github.com/google/wire"
	"github.com/go-orb/go-orb/types"
	"github.com/go-orb/go-orb/log"
	"github.com/go-orb/go-orb/registry"
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
