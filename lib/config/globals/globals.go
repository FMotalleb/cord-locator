package globals

import (
	"github.com/rs/zerolog/log"
	"strings"
)

// CoreConfiguration is holds information of dns server listen port,AllowTransfer and default provider
type CoreConfiguration struct {
	Address         string   `yaml:"address"`
	AllowTransfer   []string `yaml:"allowTransfer"`
	DefaultProvider string   `yaml:"defaultProvider"`
}

// Validate will check core configurations and verify it
func (r *CoreConfiguration) Validate() bool {
	if &r.DefaultProvider == nil {
		log.Debug().Msg("default provider has to be set")
		return false
	}
	if &r.DefaultProvider == nil {
		log.Debug().Msg("default provider has to be set")
		return false
	}
	parts := strings.SplitAfter(r.Address, ":")
	if len(parts) != 2 {
		log.Debug().Msg("address given in the config file must be something like `:53`")
		return false
	}

	return true
}
