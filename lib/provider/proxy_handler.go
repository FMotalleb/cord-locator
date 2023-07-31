//go:build !test
// +build !test

package provider

import (
	"math/rand"
	"net"

	"github.com/miekg/dns"
	"github.com/rs/zerolog/log"
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
			ResponseErrorToRequest(w, req)
			return
		}
		t := new(dns.Transfer)
		c, err := t.In(req, serverAddress)
		if err != nil {
			ResponseErrorToRequest(w, req)
			return
		}
		if err = t.Out(w, req, c); err != nil {
			log.Debug().Msgf("failed to handle request: %v", err)
			ResponseErrorToRequest(w, req)
			return
		}
		return
	}

	c := &dns.Client{Net: transport}
	resp, _, err := c.Exchange(req, serverAddress)
	if err != nil {
		log.Debug().Msgf("failed to handle request: %v", err)
		ResponseErrorToRequest(w, req)
		return
	}
	err = w.WriteMsg(resp)
	if err != nil {
		log.Debug().Msgf("failed to write response: %v", err)
	}
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
	addr := p.IP[0]
	if n := len(p.IP); n > 1 {
		addr = p.IP[rand.Intn(n)]
	}
	return addr
}

func ResponseErrorToRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetRcode(r, dns.RcodeServerFailure)
	_ = w.WriteMsg(m)
}
