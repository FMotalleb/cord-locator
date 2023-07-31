package provider

import (
	"math/rand"
	"net"

	log "github.com/rs/zerolog/log"
)

type Provider struct {
	Name string   `yaml:"name"`
	Ip   []string `yaml:"ip"`
}

func (p *Provider) GetRandomIp() string {
	if len(p.Ip) == 0 {
		log.Fatal().Msgf("provider(%s) has no ip", p.Name)
	}
	addr := p.Ip[0]
	if n := len(p.Ip); n > 1 {
		addr = p.Ip[rand.Intn(n)]
	}
	return addr
}
func (p *Provider) Validate() bool {
	if len(p.Ip) == 0 {
		return false
	}
	log.Debug().Msgf("provider: %s has %d ips", p.Name, len(p.Ip))
	for _, ip := range p.Ip {
		host, port, err := net.SplitHostPort(ip)
		if err != nil || host == "" || port == "" {
			log.Fatal().Msgf("provider: %s has error with ip %s (maybe missing port number?)", p.Name, ip)
			// return false
		}
	}
	return true
}
