package newconfig

import (
	entrypoint "github.com/FMotalleb/cord-locator/lib/new_config/entry_point"
	"github.com/FMotalleb/cord-locator/lib/new_config/provider"
	"github.com/FMotalleb/cord-locator/lib/utils/iterable"
)

// ConfigData representation of yaml config file
type ConfigData struct {
	EntryPoints []entrypoint.EntryPointData `yaml:"entry_points,flow"`
	Providers   []provider.Data             `yaml:"providers,flow"`
}

// Finalize the config data and returns new config object with given configuration
func (conf ConfigData) Finalize() Config {
	return Config{
		EntryPoints: conf.getEntryPoints(),
		Providers:   conf.getProviders(),
	}
}

func (conf ConfigData) getEntryPoints() []entrypoint.EntryPoint {
	mapper := func(entry entrypoint.EntryPointData) entrypoint.EntryPoint {
		return entrypoint.EntryPoint(entry)
	}
	return iterable.Map(conf.EntryPoints, mapper)
}
func (conf ConfigData) getProviders() []provider.Provider {
	mapper := func(entry provider.Data) provider.Provider {
		return provider.Provider(entry)
	}
	return iterable.Map(conf.Providers, mapper)
}
