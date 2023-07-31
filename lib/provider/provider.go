package provider

import (
	"net"

	"github.com/rs/zerolog/log"
)

// Provider is an external dns server (currently only supports ipv4)
type Provider struct {
	Name string   `yaml:"name"`
	IP   []string `yaml:"ip"`
}

// Validate that current provider has correct settings
func (p *Provider) Validate() bool {
	if len(p.IP) == 0 {
		return false
	}
	log.Debug().Msgf("provider: %s has %d ips", p.Name, len(p.IP))
	for _, ip := range p.IP {
		host, port, err := net.SplitHostPort(ip)
		if err != nil || host == "" || port == "" {
			log.Debug().Msgf("provider: %s has error with ip %s (maybe missing port number?)", p.Name, ip)
			return false
		}
	}
	return true
}
