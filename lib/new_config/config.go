package newconfig

import (
	entrypoint "github.com/FMotalleb/cord-locator/lib/new_config/entry_point"
	"github.com/FMotalleb/cord-locator/lib/new_config/provider"
)

type configFace interface {
	GetEntryPoints() []entrypoint.EntryPoint
	GetProviders() []provider.Provider
}

// Config of the dns server
type Config struct {
	entryPoints []entrypoint.EntryPoint
	Providers   []provider.Provider
	configFace
}

func (conf Config) GetEntryPoints() []entrypoint.EntryPoint {
	return conf.entryPoints
}

func (conf Config) GetProviders() []provider.Provider {
	return conf.Providers
}
