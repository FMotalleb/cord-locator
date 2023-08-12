//go:build !test

package provider

import (
	"github.com/miekg/dns"
	"github.com/rs/zerolog/log"
)

// Handle using this provider
func (p *Provider) Handle(transport string, req *dns.Msg) *dns.Msg {
	// serverAddress := p.getRandomIP()

	for _, serverAddress := range p.IP {
		c := &dns.Client{Net: transport}
		resp, _, err := c.Exchange(req.Copy(), serverAddress)
		if err != nil {
			name := req.Question[0].Name
			log.Debug().Msgf("failed to transfer request `%s` from %s, err: %v", name, serverAddress, err)
		} else {
			log.Debug().Msgf("responding with: %v", resp)
			return resp
		}
	}
	return nil
}
func (p *Provider) HandleTransfer(req *dns.Msg, t dns.Transfer) chan *dns.Envelope {
	for _, serverAddress := range p.IP {
		c, err := t.In(req, serverAddress)
		if err != nil {
			return nil
		}
		return c
	}
	return nil
}
