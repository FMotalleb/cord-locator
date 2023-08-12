package globals

import (
	"strings"

	"github.com/rs/zerolog/log"
)

// CoreConfiguration is holds information of dns server listen port,AllowTransfer and default provider
type CoreConfiguration struct {
	Address          string   `yaml:"address"`
	AllowTransfer    []string `yaml:"allowTransfer"`
	DefaultProviders []string `yaml:"defaultProviders"`
}

// Validate will check core configurations and verify it
func (r *CoreConfiguration) Validate() bool {
	if len(r.DefaultProviders) == 0 {
		log.Warn().Msg("default providers has to be set")
		return false
	}
	parts := strings.SplitAfter(r.Address, ":")
	if len(parts) != 2 {
		log.Warn().Msg("address given in the config file must be something like `0.0.0.0:53` or `:53`")
		return false
	}

	return true
}
