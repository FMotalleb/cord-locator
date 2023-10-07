package newconfig

import (
	entrypoint "github.com/FMotalleb/cord-locator/lib/new_config/entry_point"
	"github.com/FMotalleb/cord-locator/lib/new_config/provider"
)

// Config of the dns server
type Config struct {
	EntryPoints []entrypoint.EntryPoint
	Providers   []provider.Provider
}
