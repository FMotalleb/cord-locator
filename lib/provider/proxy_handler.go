//go:build !test

package provider

import (
	"github.com/miekg/dns"
	"github.com/rs/zerolog/log"
)

// Handle using this provider
func (p *Provider) Handle(transport string, req *dns.Msg) *dns.Msg {
	// serverAddress := p.getRandomIP()
	// if isTransfer(req) {
	// 	if transport != "tcp" {
	// 		return nil
	// 	}
	// 	t := new(dns.Transfer)
	// 	c, err := t.In(req, serverAddress)
	// 	if err != nil {
	// 		ResponseErrorToRequest(w, req)
	// 		return nil
	// 	}
	// 	if err = t.Out(w, req, c); err != nil {
	// 		log.Debug().Msgf("failed to handle request: %v", err)
	// 		ResponseErrorToRequest(w, req)
	// 		return
	// 	}
	// 	return
	// }

	for _, serverAddress := range p.IP {
		c := &dns.Client{Net: transport}
		resp, _, err := c.Exchange(req.Copy(), serverAddress)
		if err != nil {
			name := req.Question[0].Name
			log.Debug().Msgf("failed to transfer request `%s` from %s, err: %v", name, serverAddress, err)
		} else {
			log.Trace().Msgf("responding with: %v", resp)
			return resp
		}
	}
	return nil
}
