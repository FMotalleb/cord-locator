package provider

import (
	"math/rand"
	"net"

	"github.com/miekg/dns"
	log "github.com/rs/zerolog/log"
)

// HandleRequest using this provider
func (p *Provider) HandleRequest(w dns.ResponseWriter, req *dns.Msg) {
	transport := "udp"
	serverAddress := p.getRandomIP()
	if _, ok := w.RemoteAddr().(*net.TCPAddr); ok {
		transport = "tcp"
	}
	if isTransfer(req) {
		if transport != "tcp" {
			dns.HandleFailed(w, req)
			return
		}
		t := new(dns.Transfer)
		c, err := t.In(req, serverAddress)
		if err != nil {
			dns.HandleFailed(w, req)
			return
		}
		if err = t.Out(w, req, c); err != nil {
			log.Debug().Msgf("failed to handle request: %v", err)
			dns.HandleFailed(w, req)
			return
		}
		return
	}

	c := &dns.Client{Net: transport}
	resp, _, err := c.Exchange(req, serverAddress)
	if err != nil {
		log.Debug().Msgf("failed to handle request: %v", err)
		dns.HandleFailed(w, req)
		return
	}
	w.WriteMsg(resp)
}

func isTransfer(req *dns.Msg) bool {
	for _, q := range req.Question {
		switch q.Qtype {
		case dns.TypeIXFR, dns.TypeAXFR:
			return true
		}
	}
	return false
}

func (p *Provider) getRandomIP() string {
	if len(p.IP) == 0 {
		log.Fatal().Msgf("provider(%s) has no ip", p.Name)
	}
	addr := p.IP[0]
	if n := len(p.IP); n > 1 {
		addr = p.IP[rand.Intn(n)]
	}
	return addr
}
