//go:build !test

package config

import (
	"bytes"
	"net"
	"strings"
	"text/template"

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
		if rule.IsBlocked {
			provider.ResponseErrorToRequest(w, req)
			return
		}
		// TODO: fix printing resolver pointer
		log.Debug().Msgf("found rule for %s using findProvider: %v", lcName, rule.Resolver)

		raw := rule.GetRaw(dns.TypeToString[req.Question[0].Qtype])
		if raw != nil {
			mapper := make(map[string]string, 0)
			mapper["address"] = lcName
			msg, err := dns.NewRR(formatString(*raw, mapper))

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
				findProvider.Handle(w, req)
				return
			}
			log.Debug().Msgf("requested findProvider was missing please add `%v` to providers in config file", rule.Resolver)
			panic("default findProvider is missing")
		}

	}

	p := c.getDefaultProvider()
	p.Handle(w, req)
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

func formatString(text string, hashmap map[string]string) string {

	tmpl, err := template.New("Mapper").Parse(text)
	if err != nil {
		panic(err)
	}
	writer := bytes.NewBuffer(nil)
	err = tmpl.Execute(writer, hashmap)
	if err != nil {
		panic(err)
	}
	return writer.String()
}
