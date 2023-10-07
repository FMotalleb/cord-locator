package newconfig

import (
	entrypoint "github.com/FMotalleb/cord-locator/lib/new_config/entry_point"
	"github.com/FMotalleb/cord-locator/lib/new_config/provider"
	"github.com/FMotalleb/cord-locator/lib/utils/iterable"
)

type Config struct {
	EntryPoints []entrypoint.EntryPointData `yaml:"entry_points,flow"`
	Providers   []provider.ProviderData     `yaml:"providers,flow"`
}

func (conf Config) GetEntryPoints() []entrypoint.EntryPoint {
	mapper := func(entry entrypoint.EntryPointData) entrypoint.EntryPoint {
		return entrypoint.EntryPoint(entry)
	}
	return iterable.Map(conf.EntryPoints, mapper)

}
