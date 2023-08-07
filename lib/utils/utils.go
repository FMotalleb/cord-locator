//go:build !test

package utils

import (
	"bytes"
	"html/template"
	"net"
	"strings"

	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/rule"

	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/config"
	"github.com/miekg/dns"
	"github.com/rs/zerolog/log"
)

// HandleRequest checks if the request is allowed, then it finds the appropriate handler for the request and calls it.
// The handler returns a response, which is then written to the ResponseWriter.
func HandleRequest(c config.Config, w dns.ResponseWriter, req *dns.Msg) {
	log.Debug().Msgf("received request to find `%v`", req.Question)
	defer recoverDNSResponse(w, req)
	if !allowed(c.Global.AllowTransfer, w, req) {
		log.Panic().Msgf("received request is not allowed")
	}
	if len(req.Question) == 0 {
		log.Panic().Msgf("received request has no question")
	}
	requestHostname := req.Question[0].Name
	// log.Debug().Msgf("received request to find `%s`", requestHostname)
	r := c.FindRuleFor(requestHostname)

	resolvers := c.GetDefaultProvider()
	if r != nil {
		switch {
		case r.IsBlocked:
			log.Debug().Msgf("blocking `%s`", requestHostname)
			log.Panic().Msgf("blocked request")
			return
		case r.Resolver != nil:
			resolvers = c.FindProviders(r.Resolver)
			log.Debug().Msgf("handler found for `%s`, will use %v, request: %v", requestHostname, resolvers, UnNil(r.ResolverParams, requestHostname))
		case r.Raw != nil:

			if handleRawResponse(requestHostname, r, req, w) {
				log.Trace().Msgf("handled request for %s using raw response", requestHostname)
				return
			}
			log.Trace().Msgf("cannot handle request for %s using raw response", requestHostname)
		}
	} else {
		log.Debug().Msgf("no rule found for `%s`, will use default handler", requestHostname)
	}

	transport := "udp"
	if _, ok := w.RemoteAddr().(*net.TCPAddr); ok {
		transport = "tcp"
	}
	if r != nil && r.ResolverParams != nil {
		changeRequestAddress(req, *r.ResolverParams)
	}
	var resp *dns.Msg
	for _, resolver := range resolvers {
		resp = resolver.Handle(transport, req)
		if resp != nil {
			break
		}
	}

	if r != nil && r.ResolverParams != nil {
		changeResponseAddress(resp, requestHostname)
	}
	_ = w.WriteMsg(resp)
}

func handleRawResponse(requestHostname string, r *rule.Rule, req *dns.Msg, w dns.ResponseWriter) bool {
	mapper := make(map[string]string, 0)
	mapper["address"] = requestHostname
	raw := r.GetRaw(dns.TypeToString[req.Question[0].Qtype])
	if raw == nil {
		log.Error().Msgf("%s not supported in the config, continue using default handler", dns.TypeToString[req.Question[0].Qtype])
		return false
	}
	msg, err := dns.NewRR(formatString(*raw, mapper))
	if err != nil {
		log.Debug().Msgf("cannot parse raw response: %v", err)
		return false
	}
	if msg != nil {
		result := make([]dns.RR, 0)
		result = append(result, msg)
		req.Answer = result
		log.Info().Msgf("cannot parse raw response: %v", req)
		_ = w.WriteMsg(req)
		return true
	}
	return false
}
func recoverDNSResponse(w dns.ResponseWriter, req *dns.Msg) {
	err := recover()
	if err != nil {
		log.Error().Msgf("Recovering from: %v", err)
		reject(w, req)
	}
}
func changeRequestAddress(req *dns.Msg, newAddress string) *dns.Msg {
	req.Question[0].Name = newAddress
	return req
}
func changeResponseAddress(req *dns.Msg, newAddress string) *dns.Msg {
	req.Question[0].Name = newAddress
	ans := req.Answer[0]
	ansStr := strings.Replace(ans.String(), ans.Header().Name, newAddress, 1)
	result, err := dns.NewRR(ansStr)
	req.Answer[0] = result
	if err != nil {
		log.Panic().Msgf("faced an error when tried to change answer \nfrom:%v\nto:%v\nerror:%v", req.Answer[0], ansStr, err)
	}
	return req
}
func reject(w dns.ResponseWriter, r *dns.Msg) {
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
		log.Panic().Msgf("failed to use parse template: %s\nerror: %v", text, err)
	}
	writer := bytes.NewBuffer(nil)
	err = tmpl.Execute(writer, hashmap)
	if err != nil {
		log.Panic().Msgf("failed to use template for %s\nerror: %v\nhashmap:%v", text, err, hashmap)
	}
	return writer.String()
}

// FindFirst non-null value in items
//	func FindFirst[T any](items ...*T) (t *T) {
//		for _, item := range items {
//			if item != nil {
//				return item
//			}
//		}
//		return nil
//	}

// UnNil will check if value is not nil it will return value but if it was nil returns defaultValue
func UnNil[T any](value *T, defaultValue T) T {
	if value != nil {
		return *value
	}
	return defaultValue
}
