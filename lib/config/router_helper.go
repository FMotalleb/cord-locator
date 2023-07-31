package config

import (
	"net"
	"strings"

	"github.com/miekg/dns"
	"github.com/rs/zerolog/log"
)

func (c Config) Route(w dns.ResponseWriter, req *dns.Msg) {

	if len(req.Question) == 0 || !c.allowed(w, req) {
		dns.HandleFailed(w, req)
		return
	}
	lcName := strings.ToLower(req.Question[0].Name)
	log.Debug().Msgf("received request to find `%s`", lcName)
	rule := c.FindRuleFor(lcName)
	if rule != nil {
		log.Debug().Msgf("found rule for %s using provider: %s", lcName, rule.Resolver)
		provider := c.FindProvider(rule.Resolver)
		if provider != nil {
			provider.HandleRequest(w, req)
			return
		} else {
			log.Fatal().Msgf("requested provider was missing please add `%s` to providers in config file", rule.Resolver)
		}
	}

	provider := c.GetDefaultProvider()
	provider.HandleRequest(w, req)
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

func (c *Config) allowed(w dns.ResponseWriter, req *dns.Msg) bool {
	if !isTransfer(req) {
		return true
	}
	remote, _, _ := net.SplitHostPort(w.RemoteAddr().String())
	transferIPs := c.Global.AllowTransfer

	for _, ip := range transferIPs {
		if ip == remote {
			return true
		}
	}
	return false
}
