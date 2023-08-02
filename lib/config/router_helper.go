//go:build !test

package config

import (
	"net"
	"strings"

	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/provider"

	"github.com/miekg/dns"
	"github.com/rs/zerolog/log"
)

// Route will redirect dns request to resolvable provider
func (c *Config) Route(w dns.ResponseWriter, req *dns.Msg) {
	if len(req.Question) == 0 || !c.allowed(w, req) {
		provider.ResponseErrorToRequest(w, req)
		return
	}
	lcName := strings.ToLower(req.Question[0].Name)
	log.Debug().Msgf("received request to find `%s`", lcName)
	rule := c.findRuleFor(lcName)
	if rule != nil {
		// TODO: fix printing resolver pointer
		log.Debug().Msgf("found rule for %s using findProvider: %v", lcName, rule.Resolver)

		raw := rule.GetRaw(dns.TypeToString[req.Question[0].Qtype])
		if raw != nil {
			msg, err := dns.NewRR(*raw)
			if err != nil {
				log.Debug().Msgf("cannot parse raw response: %v", err)
			}
			if msg != nil {
				result := make([]dns.RR, 0)
				result = append(result, msg)
				req.Answer = result
				log.Info().Msgf("cannot parse raw response: %v", req)
				_ = w.WriteMsg(req)
				return
			}
		}
		if rule.Resolver != nil {
			findProvider := c.findProvider(*rule.Resolver)
			if findProvider != nil {
				findProvider.HandleRequest(w, req)
				return
			}
			log.Debug().Msgf("requested findProvider was missing please add `%v` to providers in config file", rule.Resolver)
			panic("default findProvider is missing")
		}

	}

	p := c.getDefaultProvider()
	p.HandleRequest(w, req)
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
