package newconfig

import (
	entrypoint "github.com/FMotalleb/cord-locator/lib/new_config/entry_point"
	"github.com/FMotalleb/cord-locator/lib/new_config/provider"
	"github.com/FMotalleb/cord-locator/lib/utils/iterable"
)

// Data representation of yaml config file
type Data struct {
	EntryPoints []entrypoint.Data `yaml:"entry_points,flow"`
	Providers   []provider.Data   `yaml:"providers,flow"`
}

// BuildConfig the config data and returns new config object with given configuration
func (conf Data) BuildConfig() Config {
	return Config{
		entryPoints: conf.getEntryPoints(),
		Providers:   conf.getProviders(),
	}
}

func (conf Data) getEntryPoints() []entrypoint.EntryPoint {
	mapper := func(entry entrypoint.Data) entrypoint.EntryPoint {
		err := entry.Validate()
		if err != nil {
			panic(err)
		}
		return entrypoint.EntryPoint(entry)
	}
	return iterable.Map(conf.EntryPoints, mapper)
}
func (conf Data) getProviders() []provider.Provider {
	mapper := func(entry provider.Data) provider.Provider {
		err := entry.Validate()
		if err != nil {
			panic(err)
		}
		return provider.Provider(entry)
	}
	return iterable.Map(conf.Providers, mapper)
}
