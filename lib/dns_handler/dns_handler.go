package dnsHandler

import "github.com/miekg/dns"

func HandleRequest(w dns.ResponseWriter, r *dns.Msg) {

}
func ResponseErrorToRequest(w dns.ResponseWriter, r *dns.Msg) {
	msg := MakeErrorMessage(r)
	_ = w.WriteMsg(msg)
}

func MakeErrorMessage(r *dns.Msg) *dns.Msg {
	m := new(dns.Msg)
	m.SetRcode(r, dns.RcodeServerFailure)
	return m
}
