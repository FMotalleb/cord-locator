package iResolver

import "github.com/miekg/dns"

type IResolver interface {
	Resolve(address string) dns.Msg
}
