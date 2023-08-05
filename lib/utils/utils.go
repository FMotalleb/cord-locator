//go:build !test

package utils

import (
	"bytes"
	"html/template"
	"net"

	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/config"
	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/provider"
	"github.com/miekg/dns"
	"github.com/rs/zerolog/log"
)

// HandleRequest checks if the request is allowed, then it finds the appropriate handler for the request and calls it.
// The handler returns a response, which is then written to the ResponseWriter.
func HandleRequest(c config.Config, w dns.ResponseWriter, req *dns.Msg) {
	log.Debug().Msgf("received request to find `%v`", req.Question)
	if len(req.Question) == 0 || !allowed(c.Global.AllowTransfer, w, req) {
		responseErrorToRequest(w, req)
		return
	}
	requestHostname := req.Question[0].Name
	log.Debug().Msgf("received request to find `%s`", requestHostname)
	r := c.FindRuleFor(requestHostname)
	dnsProvider := provider.Provider{}
	switch {
	case r == nil:
		dnsProvider = *c.GetDefaultProvider()
	case r.Resolver != nil:
		dnsProvider = *c.FindProvider(*r.Resolver)
	case r.Raw != nil:
		mapper := make(map[string]string, 0)
		mapper["address"] = requestHostname
		raw := r.GetRaw(dns.TypeToString[req.Question[0].Qtype])
		if raw == nil {
			log.Error().Msgf("%s not supported in the config", dns.TypeToString[req.Question[0].Qtype])
			responseErrorToRequest(w, req)
			return
		}
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
	log.Debug().Msgf("no handler found for `%s`, will use default handler", requestHostname)
	transport := "udp"
	if _, ok := w.RemoteAddr().(*net.TCPAddr); ok {
		transport = "tcp"
	}
	resp := dnsProvider.Handle(transport, req)
	_ = w.WriteMsg(resp)
}

func responseErrorToRequest(w dns.ResponseWriter, r *dns.Msg) {
	msg := makeErrorMessage(r)
	_ = w.WriteMsg(msg)
}

func makeErrorMessage(r *dns.Msg) *dns.Msg {
	m := new(dns.Msg)
	m.SetRcode(r, dns.RcodeServerFailure)
	return m
}

func allowed(transferIPs []string, w dns.ResponseWriter, req *dns.Msg) bool {
	if !isTransfer(req) {
		return true
	}
	remote, _, _ := net.SplitHostPort(w.RemoteAddr().String())

	for _, ip := range transferIPs {
		if ip == remote {
			return true
		}
	}
	return false
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

// FindFirst non-null value in items
func FindFirst[T any](items ...*T) (t *T) {
	for _, item := range items {
		if item != nil {
			return item
		}
	}
	return nil
}
